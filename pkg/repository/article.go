package repository

import (
	"log"
	"time"

	"github.com/gelerum/artpaper/pkg/model"
)

func (r *Repository) CreateArticle(articlename string, title string, body string, creationDate string, username string) (err error) {
	_, err = r.database.Exec("INSERT INTO articles(articlename, title, body, creation_date, username) VALUES ($1, $2, $3, $4, $5);", articlename, title, body, creationDate, username)
	if err != nil {
		log.Printf("Error occurred on article creation from database: %s", err)
	}
	return
}

func (r *Repository) GetArticle(articlename string) (title string, body string, creationDate time.Time, username string, err error) {
	row := r.database.QueryRow("SELECT title, body, creation_date, username FROM articles WHERE articlename = $1;", articlename)
	err = row.Scan(&title, &body, &creationDate, &username)
	if err != nil {
		log.Printf("Error occurred on article selection from database: %s", err)
	}
	return
}

func (r *Repository) UpdateArticle(articlename string, newArticlename, newTitle string, newBody string) (err error) {
	_, err = r.database.Exec("UPDATE articles SET articlename = $2, title = $3, body = $4 WHERE articlename = $1;", articlename, newArticlename, newTitle, newBody)
	if err != nil {
		log.Printf("Error occurred on article updation from database: %s", err)
	}
	return
}

func (r *Repository) DeleteArticle(articlename string) (err error) {
	_, err = r.database.Exec("DELETE FROM articles WHERE articlename = $1;", articlename)
	if err != nil {
		log.Printf("Error occurred on article deletion from database: %s", err)
	}
	return
}

func (r *Repository) DeleteAllArticles(username string) (err error) {
	_, err = r.database.Exec("DELETE FROM articles WHERE username = $1;", username)
	if err != nil {
		log.Printf("Error occurred on all articles deletion from database: %s", err)
	}
	return
}

func (r *Repository) GetArticles(title string, username string, from string, to string, quantity int) (articles []model.Article, err error) {
	title += "%"
	username += "%"
	rows, err := r.database.Query("SELECT articlename, title, body, creation_date, username FROM articles WHERE title LIKE $1 AND username LIKE $2 AND creation_date >= $3 AND creation_date <= $4 LIMIT $5;", title, username, from, to, quantity)
	if err != nil {
		log.Printf("Error occurred on articles selection from database: %s", err)
		return
	}
	for rows.Next() {
		var (
			articlename  string
			title        string
			body         string
			creationDate time.Time
			username     string
		)
		err = rows.Scan(&articlename, &title, &body, &creationDate, &username)
		if err != nil {
			log.Println(err)
			return
		}
		article := model.Article{
			Articlename:  articlename,
			Title:        title,
			Body:         body,
			CreationDate: creationDate,
			Username:     username,
		}
		articles = append(articles, article)
	}
	return
}

func (r *Repository) ArticlesNameCount(title string) (count int, err error) {
	row := r.database.QueryRow("SELECT COUNT(*) FROM articles WHERE title = $1;", title)
	err = row.Scan(&count)
	if err != nil {
		log.Printf("Error occurred on articles selection from database: %s", err)
	}
	return
}
