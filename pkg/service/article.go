package service

import (
	"strconv"
	"strings"
	"time"

	"github.com/gelerum/artpaper/pkg/model"
)

func (s *Service) CreateArticle(title string, body string, username string) (articlename string, err error) {
	articlename, err = s.generateArticlename(title)
	if err != nil {
		return
	}
	err = s.database.CreateArticle(articlename, title, body, time.Now().Format("2006-01-02"), username)
	return
}

func (s *Service) GetArticle(articlename string) (title string, body string, creationDate time.Time, username string, err error) {
	title, body, creationDate, username, err = s.database.GetArticle(articlename)
	if err != nil {
		return
	}
	return
}

func (s *Service) UpdateArticle(articlename string, newTitle string, newBody string) (err error) {
	newArticlename, err := s.generateArticlename(newTitle)
	if err != nil {
		return
	}
	err = s.database.UpdateArticle(articlename, newArticlename, newTitle, newBody)
	return
}

func (s *Service) DeleteArticle(articlename string) (err error) {
	err = s.database.DeleteArticle(articlename)
	return
}

func (s *Service) GetArticles(title string, username string, from string, to string, quantity int) (articles []model.Article, err error) {
	articles, err = s.database.GetArticles(title, username, from, to, quantity)
	return
}

func (s *Service) generateArticlename(title string) (articlename string, err error) {
	count, err := s.database.ArticlesNameCount(title)
	if err != nil {
		return
	}
	articlename = strings.ReplaceAll(title, " ", "-")
	if count != 0 {
		articlename += "-" + strconv.Itoa(count)
	}
	return
}
