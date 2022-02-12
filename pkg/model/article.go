package model

import "time"

type Article struct {
	Articlename  string    `json:"articlename"`
	Title        string    `json:"title"`
	Body         string    `json:"body"`
	CreationDate time.Time `json:"creationDate"`
	Username     string    `json:"username"`
}
