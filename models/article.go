package models

import "github.com/MohamedSaidCS/web-scraper-api/db"

type Article struct {
	ID       int64  `json:"id"`
	Title    string `json:"title" binding:"required"`
	Link     string `json:"link" binding:"required"`
	Timesamp string `json:"timestamp"`
}

func (a *Article) Create() error {
	query := "INSERT INTO articles (title, link) VALUES ($1, $2) RETURNING id, timestamp"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRow(a.Title, a.Link).Scan(&a.ID, &a.Timesamp)
	if err != nil {
		return err
	}

	return nil
}

func GetArticles() ([]Article, error) {
	query := "SELECT * FROM articles"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	articles := []Article{}

	for rows.Next() {
		var article Article
		err := rows.Scan(&article.ID, &article.Title, &article.Link, &article.Timesamp)
		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	return articles, nil
}

func GetArticleByID(id int64) (Article, error) {
	query := "SELECT * FROM articles WHERE id=$1"
	row := db.DB.QueryRow(query, id)

	var article Article
	err := row.Scan(&article.ID, &article.Title, &article.Link, &article.Timesamp)
	if err != nil {
		return Article{}, err
	}

	return article, nil
}

func (a *Article) Update() error {
	query := "UPDATE articles SET title=$1, link=$2 WHERE id=$3"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(a.Title, a.Link, a.ID)
	if err != nil {
		return err
	}

	return nil
}

func (a *Article) Delete() error {
	query := "DELETE FROM articles WHERE id=$1"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(a.ID)
	if err != nil {
		return err
	}

	return nil
}
