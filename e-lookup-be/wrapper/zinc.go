package wrapper

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	m "elookup/wrapper/model"
)

type _QueryHits struct {
	TotalHits int       `json:"totalHits"`
	Hits      []m.Email `json:"hits"`
}

func SearchByWord(term string, from int, maxResult int) _QueryHits {
	url := getZincSearchServerURL() + "enron-index" + "/_search"

	queryBody := m.ZincSearchQueryRequest{
		SearchType: "matchphrase",
		Query: m.ZincQuery{
			Term: term,
		},
		From:       from * maxResult,
		MaxResults: maxResult,
		Fields:     []string{},
		Highlight: m.ZincHighlight{
			Fields: map[string]any{
				"content": map[string]any{},
			},
		},
	}

	rqbody, _ := json.Marshal(queryBody)
	strRqBody := string(rqbody)
	req, err := http.NewRequest("POST", url, strings.NewReader(strRqBody))
	log.Println("Posting to Zinc. Query body:", strRqBody)
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
	log.Printf("Zinc response code: %v", res.StatusCode)

	queryResponse, _ := io.ReadAll(res.Body)
	data := m.ZincSearchQueryResponse{}
	json.Unmarshal([]byte(queryResponse), &data)
	return MapResponseToMails(data)
}

func MapResponseToMails(zincResponse m.ZincSearchQueryResponse) _QueryHits {
	mails := []m.Email{}
	for _, item := range zincResponse.Hits.ActualEmails {
		mails = append(mails, m.Email{
			Id:        item.Id,
			Date:      item.Source.Date,
			From:      item.Source.From,
			To:        item.Source.To,
			Subject:   item.Source.Subject,
			Content:   item.Source.Content,
			Highlight: item.Highlight.Content,
		})
	}
	return _QueryHits{
		TotalHits: zincResponse.Hits.Total.Value,
		Hits:      mails,
	}
}

func getZincSearchServerURL() string {
	return "http://localhost:4080/api/"
	// return os.Getenv("ZINC_SEARCH_SERVER_URL")
}
func getZincSearchUser() string {
	return "admin"
	// return os.Getenv("ZINC_SEARCH_USER")
}
func getZincSearchPassword() string {
	return "Complexpass#123"
	// return os.Getenv("ZINC_SEARCH_PASSWORD")
}
