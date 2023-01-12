package wrapper

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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

type EmailHits struct {
	Total  int     `json:"total"`
	Emails []Email `json:"emails"`
}

func SearchByWord(term string, from int, maxResult int) EmailHits {
	url := getZincSearchServerURL() + "enron-index" + "/_search"

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
			MessageId: item.Id,
			Date:      item.ItemSource.Date,
			From:      item.ItemSource.From,
			To:        item.ItemSource.To,
			Subject:   item.ItemSource.Subject,
			Content:   item.ItemSource.Content,
		})
	}
	return EmailHits{
		Total:  hitsValues.Hits.Total.Value,
		Emails: mails,
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
