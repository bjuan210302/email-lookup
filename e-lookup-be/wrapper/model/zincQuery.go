package model

// REQUEST
type ZincSearchQueryRequest struct {
	SearchType string        `json:"search_type"`
	Query      ZincQuery     `json:"query"`
	From       int           `json:"from"`
	MaxResults int           `json:"max_results"`
	Fields     []string      `json:"_source"`
	Highlight  ZincHighlight `json:"highlight"`
}

type ZincQuery struct {
	Term string `json:"term"`
}

type ZincHighlight struct {
	Fields map[string]any `json:"fields"`
}

// RESPOSE
type ZincSearchQueryResponse struct {
	MaxScore float32 `json:"max_score"`
	Hits     struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`

		ActualEmails []_EmailHit `json:"hits"`
	} `json:"hits"`
}

type _EmailHit struct {
	Id        string `json:"_id"`
	Source    Email  `json:"_source"`
	Highlight struct {
		Content []string `json:"content"`
	} `json:"highlight"`
}
