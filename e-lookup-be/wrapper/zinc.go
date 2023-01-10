package wrapper

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type QueryOptions struct {
	Term string `json:"term"`
}

// Minimal query
type ZincSearchQueryRequest struct {
	SearchType string       `json:"search_type"`
	Query      QueryOptions `json:"query"`
	From       int          `json:"from,omitempty"`
	MaxResults int          `json:"max_results,omitempty"`
	Source     []string     `json:"_source,omitempty"`
}

type ZincSearchQueryResponse struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		MaxScore    float32       `json:"max_score"`
		HitsSummary []HitsSummary `json:"hits"`
	} `json:"hits"`
}

type HitsSummary struct {
	Id         string `json:"_id"`
	ItemSource Email  `json:"_source"`
}

// Minimal
type Email struct {
	Id      string   `json:"id"`
	From    string   `json:"sender"`
	To      []string `json:"recipient"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
}

type EmailHits struct {
	Total  int     `json:"total"`
	Emails []Email `json:"emails"`
}

func SearchByWord(term string, from int, maxResult int) EmailHits {
	url := getZincSearchServerURL() + "enron" + "/_search"

	queryBody := ZincSearchQueryRequest{
		SearchType: "matchphrase",
		Query: QueryOptions{
			Term: term,
		},
		From:       from * maxResult,
		MaxResults: maxResult,
		Source:     []string{},
	}

	rqbody, _ := json.Marshal(queryBody)
	req, err := http.NewRequest("POST", url, strings.NewReader(string(rqbody)))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(getZincSearchUser(), getZincSearchPassword())
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	log.Println(res.StatusCode)
	queryResponse, _ := io.ReadAll(res.Body)
	data := ZincSearchQueryResponse{}
	json.Unmarshal([]byte(queryResponse), &data)
	return MapResponseToMails(data)
}

func MapResponseToMails(hitsValues ZincSearchQueryResponse) EmailHits {
	mails := []Email{}
	for _, item := range hitsValues.Hits.HitsSummary {
		mails = append(mails, Email{
			Id:      item.Id,
			Subject: item.ItemSource.Subject,
			From:    item.ItemSource.From,
			To:      item.ItemSource.To,
		})
	}
	return EmailHits{
		Total:  hitsValues.Hits.Total.Value,
		Emails: mails,
	}
}

func getZincSearchServerURL() string {
	return os.Getenv("ZINC_SEARCH_SERVER_URL")
}
func getZincSearchUser() string {
	return os.Getenv("ZINC_SEARCH_USER")
}
func getZincSearchPassword() string {
	return os.Getenv("ZINC_SEARCH_PASSWORD")
}
