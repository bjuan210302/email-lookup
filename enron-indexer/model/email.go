package model

type Email struct {
	MessageId string   `json:"id"`
	From      string   `json:"from"`
	To        []string `json:"to"`
	Subject   string   `json:"subject"`
	Content   string   `json:"content"`
	Date      string   `json:"date"`
}
