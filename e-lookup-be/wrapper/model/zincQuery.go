package model

// REQUEST
type ZincSearchQueryRequest struct {
	SearchType string        `json:"search_type"`
	Query      ZincQuery     `json:"query"`
	From       int           `json:"from,omitempty"`
	MaxResults int           `json:"max_results,omitempty"`
	Fields     []string      `json:"_source,omitempty"`
	Highlight  ZincHighlight `json:"highlight,omitempty"`
}

type ZincQuery struct {
	Term string `json:"term"`
}

type ZincHighlight struct {
	PreTags  []string `json:"pre_tags,omitempty"`
	PostTags []string `json:"post_tags,omitempty"`
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
