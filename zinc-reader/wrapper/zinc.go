package wrapper

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"zincreader/model"
)

func SaveBulk(mails []model.Email) {
	// Using bulk v2 as it is more confortable than bulk
	url := getZincSearchAPIURL() + "_bulkv2"
	name := getZincSearchMailIndexName()
	bulk := struct {
		IndexName string `json:"index"`
		Records   any    `json:"records"`
	}{
		IndexName: name,
		Records:   mails,
	}
	bulkJson, _ := json.Marshal(bulk)

	log.Print("Posting bulk to Zinc server")
	httpPOST(url, string(bulkJson))
}

// This is needed because when an index is created through bulk upload properties are NOT highlightable by default.
// It seems that it's not posible to update a property from non-highlightable to highlightable, so it's necessary
// to create the index manually with highlightable properties
func CheckAndCreateIndex() {
	log.Printf("Checking if index %s exists and is highlightable", getZincSearchMailIndexName())
	url := getZincSearchAPIURL() + getZincSearchMailIndexName() + "/_mapping"
	req, _ := makeRequestWithAuth("GET", url, "")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Index does not exist, create it
	if res.StatusCode == 400 {
		log.Printf("Index %s does not exist. Creating it...", getZincSearchMailIndexName())
		createIndex()
		return
	}

	// Index does exist, check if highlightable
	if res.StatusCode == 200 {
		log.Printf("Index %s does exist. Checking if highlightable...", getZincSearchMailIndexName())
		bodyBytes, _ := io.ReadAll(res.Body)

		type IndexInfo struct {
			Mappings struct {
				Properties map[string]struct {
					Highlightable bool `json:"highlightable"`
				} `json:"properties"`
			} `json:"mappings"`
		}

		index := make(map[string]IndexInfo)
		json.Unmarshal([]byte(bodyBytes), &index)

		if !index[getZincSearchMailIndexName()].Mappings.Properties["content"].Highlightable {
			log.Print("Index is not highlightable, deleting index and creating again...")
			deleteIndex()
			createIndex()
		}
	}

	log.Print("Index checks finalized")

}

func createIndex() {
	log.Printf("Creating index: %s", getZincSearchMailIndexName())
	url := getZincSearchAPIURL() + "index"
	indexTemplate := `{
		"name": "` + getZincSearchMailIndexName() + `",
		"storage_type": "disk",
		"shard_num": 1,
		"mappings": {
			"properties": {
				"from": {
					"type": "text",
					"index": false,
					"store": false
				},
				"to": {
					"type": "text",
					"index": false,
					"store": false
				},
				"subject": {
					"type": "text",
					"index": true,
					"store": false
				},
				"content": {
					"type": "text",
					"index": true,
					"store": false,
					"highlightable": true
				}
			}
		}
	}`
	httpPOST(url, indexTemplate)
}

func deleteIndex() {
	log.Printf("Deleting index: %s", getZincSearchMailIndexName())
	url := getZincSearchAPIURL() + "index/" + getZincSearchMailIndexName()
	req, _ := makeRequestWithAuth("DELETE", url, "")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Zinc server response code: %d", res.StatusCode)
}

func httpPOST(url string, body string) {
	req, err := makeRequestWithAuth("POST", url, body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Posting to: %s...", url)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	log.Printf("Zinc server response code: %d", res.StatusCode)
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Zinc server response body: %s", string(resBody))
}

func makeRequestWithAuth(method string, url string, body string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	zincUser := os.Getenv("ZINC_SEARCH_USER")
	zincPass := os.Getenv("ZINC_SEARCH_PASSWORD")
	req.SetBasicAuth(zincUser, zincPass)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func getZincSearchAPIURL() string {
	return os.Getenv("ZINC_SEARCH_SERVER_URL")
}
func getZincSearchMailIndexName() string {
	return os.Getenv("ZINC_SEARCH_INDEX_NAME")
}
