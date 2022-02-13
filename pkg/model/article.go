package model

import "time"

type Article struct {
	Articlename  string    `json:"articlename"`
	Title        string    `json:"title"`
	Body         string    `json:"body"`
	CreationDate time.Time `json:"creationDate"`
	Username     string    `json:"username"`
}

type Articlename struct {
	Articlename string `json:"articlename"`
}

type CreateArticle struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	Username string `json:"username"`
}

type GetArticle struct {
	Title        string `json:"title"`
	Body         string `json:"body"`
	CreationDate string `json:"creationDate"`
	Username     string `json:"username"`
}

type UpdateArticle struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
