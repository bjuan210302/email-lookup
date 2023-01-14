package model

type Email struct {
	Id        string   `json:"_id"`
	Date      string   `json:"date"`
	From      string   `json:"from"`
	To        []string `json:"to"`
	Subject   string   `json:"subject"`
	Content   string   `json:"content"`
	Highlight []string `json:"highlight"`
}
