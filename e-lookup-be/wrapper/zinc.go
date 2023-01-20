package wrapper

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	m "elookup/wrapper/model"
)

type _QueryHits struct {
	TotalHits int       `json:"totalHits"`
	Hits      []m.Email `json:"hits"`
}

func SearchByWord(term string, from int, maxResult int, indexName string, auth string) (*_QueryHits, error) {
	url := getZincSearchServerURL() + indexName + "/_search"

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
	rqbody, err := json.Marshal(queryBody)
	if err != nil {
		return nil, err
	}
	strRqBody := string(rqbody)

	log.Println("SearchByWord: Query link:", url)
	log.Println("SearchByWord: Query body:", strRqBody)
	req, err := http.NewRequest("POST", url, strings.NewReader(strRqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", auth)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 200 {
		data := m.ZincSearchQueryResponse{}
		json.Unmarshal([]byte(resBody), &data)
		return MapResponseToMails(data), nil
	} else {
		return nil, m.NewRequestError(res.StatusCode, string(resBody))
	}
}

func MapResponseToMails(zincResponse m.ZincSearchQueryResponse) *_QueryHits {
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
	return &_QueryHits{
		TotalHits: zincResponse.Hits.Total.Value,
		Hits:      mails,
	}
}

func GetIndexNamesList(auth string) ([]string, error) {
	url := getZincSearchServerURL() + "index_name"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", auth)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	indexNames := []string{}
	if res.StatusCode == 200 {
		json.Unmarshal(resBody, &indexNames)
	} else {
		return nil, m.NewRequestError(res.StatusCode, string(resBody))
	}

	return indexNames, nil
}

func getZincSearchServerURL() string {
	return os.Getenv("ZINC_SEARCH_SERVER_URL")
}
