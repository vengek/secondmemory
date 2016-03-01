package models

type Record struct {
	Id 	int        `json:"id"`
	Type 	int        `json:"type"`
	Text 	string     `json:"text"`
	Href 	string     `json:"href"`
}
