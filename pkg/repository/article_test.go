package repository

import (
	"testing"
	"time"

	"github.com/gelerum/artpaper/pkg/model"
	"github.com/jackc/pgx"
)

func TestCreateArticle(t *testing.T) {
	// before
	r, err := NewRepository()
	if err != nil {
		t.Fatal(err)
	}
	insertedUsername := "testusername"
	insertedName := "testname"
	insertedPassword := "testpassword"
	_, err = r.database.Exec("INSERT INTO users(username, name, password) VALUES ($1, $2, $3)", insertedUsername, insertedName, insertedPassword)
	if err != nil {
		t.Error(err)
	}
	// testing
	insertedArticlename := "testarticlename"
	insertedTitle := "testtitle"
	insertedBody := "testbody"
	insertedCreationDate := time.Now().Format("2006-01-02")
	err = r.CreateArticle(insertedArticlename, insertedTitle, insertedBody, insertedCreationDate, insertedUsername)
	if err != nil {
		t.Error(err)
	}

	var (
		receivedArticlename  string
		receivedTitle        string
		receivedBody         string
		receivedCreationDate time.Time
	)
	row := r.database.QueryRow("SELECT articlename, title, body, creation_date FROM articles WHERE username = $1;", insertedUsername)
	err = row.Scan(&receivedArticlename, &receivedTitle, &receivedBody, &receivedCreationDate)
	if err != nil {
		t.Error(err)
	}
	if insertedArticlename != receivedArticlename {
		t.Error("Doesn't insert properly", insertedArticlename, "!=", receivedArticlename)
	}
	if insertedTitle != receivedTitle {
		t.Error("Doesn't insert properly", insertedTitle, "!=", receivedTitle)
	}
	if insertedBody != receivedBody {
		t.Error("Doesn't insert properly", insertedBody, "!=", receivedBody)
	}
	if insertedCreationDate != receivedCreationDate.Format("2006-01-02") {
		t.Error("Doesn't insert properly", insertedCreationDate, "!=", receivedCreationDate)
	}
	// after
	_, err = r.database.Exec("DELETE FROM articles *;")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.database.Exec("DELETE FROM users *;")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetArticle(t *testing.T) {
	// before
	r, err := NewRepository()
	if err != nil {
		t.Fatal(err)
	}
	insertedUsername := "testusername"
	insertedName := "testname"
	insertedPassword := "testpassword"
	_, err = r.database.Exec("INSERT INTO users(username, name, password) VALUES ($1, $2, $3)", insertedUsername, insertedName, insertedPassword)
	if err != nil {
		t.Error(err)
	}
	insertedArticlename := "testarticlename"
	insertedTitle := "testtitle"
	insertedBody := "testbody"
	insertedCreationDate := time.Now().Format("2006-01-02")
	_, err = r.database.Exec("INSERT INTO articles(articlename, title, body, creation_date, username) VALUES ($1, $2, $3, $4, $5);", insertedArticlename, insertedTitle, insertedBody, insertedCreationDate, insertedUsername)
	if err != nil {
		t.Error(err)
	}
	// testing
	receivedTitle, receivedBody, receivedCreationDate, receivedUsername, err := r.GetArticle(insertedArticlename)
	if err != nil {
		t.Error(err)
	}

	if insertedTitle != receivedTitle {
		t.Error("Doesn't receive properly", insertedTitle, "!=", receivedTitle)
	}
	if insertedBody != receivedBody {
		t.Error("Doesn't receive properly", insertedBody, "!=", receivedBody)
	}
	if insertedCreationDate != receivedCreationDate.Format("2006-01-02") {
		t.Error("Doesn't receive properly", insertedCreationDate, "!=", receivedCreationDate)
	}
	if insertedUsername != receivedUsername {
		t.Error("Doesn't receive properly", insertedUsername, "!=", receivedUsername)
	}
	// after
	_, err = r.database.Exec("DELETE FROM articles *;")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.database.Exec("DELETE FROM users *;")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateArticle(t *testing.T) {
	// before
	r, err := NewRepository()
	if err != nil {
		t.Fatal(err)
	}
	insertedUsername := "testusername"
	insertedName := "testname"
	insertedPassword := "testpassword"
	_, err = r.database.Exec("INSERT INTO users(username, name, password) VALUES ($1, $2, $3)", insertedUsername, insertedName, insertedPassword)
	if err != nil {
		t.Error(err)
	}
	insertedArticlename := "testarticlename"
	insertedTitle := "testtitle"
	insertedBody := "testbody"
	insertedCreationDate := time.Now().Format("2006-01-02")
	_, err = r.database.Exec("INSERT INTO articles(articlename, title, body, creation_date, username) VALUES ($1, $2, $3, $4, $5);", insertedArticlename, insertedTitle, insertedBody, insertedCreationDate, insertedUsername)
	if err != nil {
		t.Error(err)
	}
	// testing
	updatedArticlename := "updatedarticlename"
	updatedTitle := "updatedtitle"
	updatedBody := "updatedbody"
	err = r.UpdateArticle(insertedArticlename, updatedArticlename, updatedTitle, updatedBody)
	if err != nil {
		t.Error(err)
	}

	var (
		receivedArticlename string
		receivedTitle       string
		receivedBody        string
	)
	row := r.database.QueryRow("SELECT articlename, title, body FROM articles WHERE username = $1;", insertedUsername)
	err = row.Scan(&receivedArticlename, &receivedTitle, &receivedBody)
	if err != nil {
		t.Error(err)
	}
	if updatedArticlename != receivedArticlename {
		t.Error("Doesn't insert properly", updatedArticlename, "!=", receivedArticlename)
	}
	if updatedTitle != receivedTitle {
		t.Error("Doesn't insert properly", updatedTitle, "!=", receivedTitle)
	}
	if updatedBody != receivedBody {
		t.Error("Doesn't insert properly", updatedBody, "!=", receivedBody)
	}
	// after
	_, err = r.database.Exec("DELETE FROM articles *;")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.database.Exec("DELETE FROM users *;")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteArticle(t *testing.T) {
	// before
	r, err := NewRepository()
	if err != nil {
		t.Fatal(err)
	}
	insertedUsername := "testusername"
	insertedName := "testname"
	insertedPassword := "testpassword"
	_, err = r.database.Exec("INSERT INTO users(username, name, password) VALUES ($1, $2, $3)", insertedUsername, insertedName, insertedPassword)
	if err != nil {
		t.Error(err)
	}
	insertedArticlename := "testarticlename"
	insertedTitle := "testtitle"
	insertedBody := "testbody"
	insertedCreationDate := time.Now().Format("2006-01-02")
	_, err = r.database.Exec("INSERT INTO articles(articlename, title, body, creation_date, username) VALUES ($1, $2, $3, $4, $5);", insertedArticlename, insertedTitle, insertedBody, insertedCreationDate, insertedUsername)
	if err != nil {
		t.Error(err)
	}
	// testing
	err = r.DeleteArticle(insertedArticlename)
	if err != nil {
		t.Error(err)
	}

	row := r.database.QueryRow("SELECT * FROM articles WHERE articlename = $1;", insertedArticlename)
	err = row.Scan()
	if err != pgx.ErrNoRows {
		t.Error(err)
	}
	// after
	_, err = r.database.Exec("DELETE FROM articles *;")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.database.Exec("DELETE FROM users *;")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteAllArticles(t *testing.T) {
	// before
	r, err := NewRepository()
	if err != nil {
		t.Fatal(err)
	}
	insertedUsername := "testusername"
	insertedName := "testname"
	insertedPassword := "testpassword"
	_, err = r.database.Exec("INSERT INTO users(username, name, password) VALUES ($1, $2, $3)", insertedUsername, insertedName, insertedPassword)
	if err != nil {
		t.Error(err)
	}
	insertedTitle := "testtitle"
	insertedBody := "testbody"
	insertedCreationDate := time.Now().Format("2006-01-02")

	insertedFirstArticlename := "1testarticlename"
	_, err = r.database.Exec("INSERT INTO articles(articlename, title, body, creation_date, username) VALUES ($1, $2, $3, $4, $5);", insertedFirstArticlename, insertedTitle, insertedBody, insertedCreationDate, insertedUsername)
	if err != nil {
		t.Error(err)
	}
	insertedSecondArticlename := "2testarticlename"
	_, err = r.database.Exec("INSERT INTO articles(articlename, title, body, creation_date, username) VALUES ($1, $2, $3, $4, $5);", insertedSecondArticlename, insertedTitle, insertedBody, insertedCreationDate, insertedUsername)
	if err != nil {
		t.Error(err)
	}
	// testing
	err = r.DeleteAllArticles(insertedUsername)
	if err != nil {
		t.Error(err)
	}

	row := r.database.QueryRow("SELECT * FROM articles WHERE username = $1;", insertedUsername)
	err = row.Scan()
	if err != pgx.ErrNoRows {
		t.Error(err)
	}
	// after
	_, err = r.database.Exec("DELETE FROM articles *;")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.database.Exec("DELETE FROM users *;")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetArticles(t *testing.T) {
	// before
	r, err := NewRepository()
	if err != nil {
		t.Fatal(err)
	}
	insertedUsername := "testusername"
	insertedName := "testname"
	insertedPassword := "testpassword"
	_, err = r.database.Exec("INSERT INTO users(username, name, password) VALUES ($1, $2, $3)", insertedUsername, insertedName, insertedPassword)
	if err != nil {
		t.Error(err)
	}
	insertedTitle := "testtitle"
	insertedBody := "testbody"
	insertedCreationDate := time.Now().Format("2006-01-02")

	insertedFirstArticlename := "1testarticlename"
	_, err = r.database.Exec("INSERT INTO articles(articlename, title, body, creation_date, username) VALUES ($1, $2, $3, $4, $5);", insertedFirstArticlename, insertedTitle, insertedBody, insertedCreationDate, insertedUsername)
	if err != nil {
		t.Error(err)
	}
	insertedSecondArticlename := "2testarticlename"
	_, err = r.database.Exec("INSERT INTO articles(articlename, title, body, creation_date, username) VALUES ($1, $2, $3, $4, $5);", insertedSecondArticlename, insertedTitle, insertedBody, insertedCreationDate, insertedUsername)
	if err != nil {
		t.Error(err)
	}
	// testing
	articles, err := r.GetArticles(insertedTitle, insertedUsername, "0001-01-01", "9999-01-01", 2)
	if err != nil {
		t.Error(err)
	}
	date, err := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	if err != nil {
		t.Error(err)
	}
	firstArticle := model.Article{
		Articlename:  insertedFirstArticlename,
		Title:        insertedTitle,
		Body:         insertedBody,
		CreationDate: date,
		Username:     insertedUsername,
	}
	secondArticle := model.Article{
		Articlename:  insertedSecondArticlename,
		Title:        insertedTitle,
		Body:         insertedBody,
		CreationDate: date,
		Username:     insertedUsername,
	}
	if articles[0] != firstArticle {
		t.Error("Doesn't receive properly\n", articles[0], "!=\n", firstArticle)
	}
	if articles[1] != secondArticle {
		t.Error("Doesn't receive properly\n", articles[1], "!=\n", secondArticle)
	}
	// after
	_, err = r.database.Exec("DELETE FROM articles *;")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.database.Exec("DELETE FROM users *;")
	if err != nil {
		t.Fatal(err)
	}
}

func TestArticlesNameCount(t *testing.T) {
	// before
	r, err := NewRepository()
	if err != nil {
		t.Fatal(err)
	}
	insertedUsername := "testusername"
	insertedName := "testname"
	insertedPassword := "testpassword"
	_, err = r.database.Exec("INSERT INTO users(username, name, password) VALUES ($1, $2, $3)", insertedUsername, insertedName, insertedPassword)
	if err != nil {
		t.Error(err)
	}
	insertedTitle := "testtitle"
	insertedBody := "testbody"
	insertedCreationDate := time.Now().Format("2006-01-02")

	insertedFirstArticlename := "1testarticlename"
	_, err = r.database.Exec("INSERT INTO articles(articlename, title, body, creation_date, username) VALUES ($1, $2, $3, $4, $5);", insertedFirstArticlename, insertedTitle, insertedBody, insertedCreationDate, insertedUsername)
	if err != nil {
		t.Error(err)
	}
	insertedSecondArticlename := "2testarticlename"
	_, err = r.database.Exec("INSERT INTO articles(articlename, title, body, creation_date, username) VALUES ($1, $2, $3, $4, $5);", insertedSecondArticlename, insertedTitle, insertedBody, insertedCreationDate, insertedUsername)
	if err != nil {
		t.Error(err)
	}
	// testing
	count, err := r.ArticlesNameCount(insertedTitle)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("Doesn't count properly")
	}
	// after
	_, err = r.database.Exec("DELETE FROM articles *;")
	if err != nil {
		t.Fatal(err)
	}
	_, err = r.database.Exec("DELETE FROM users *;")
	if err != nil {
		t.Fatal(err)
	}
}
