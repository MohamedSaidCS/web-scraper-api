package models

import "github.com/MohamedSaidCS/web-scraper-api/db"

type Article struct {
	ID    int64
	Title string `binding:"required"`
	Link  string `binding:"required"`
}

func (a *Article) Create() error {
	query := "INSERT INTO articles (title, link) VALUES (?, ?) RETURNING id"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRow(a.Title, a.Link).Scan(&a.ID)
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
		err := rows.Scan(&article.ID, &article.Title, &article.Link)
		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	return articles, nil
}

func GetArticleByID(id int64) (Article, error) {
	query := "SELECT * FROM articles WHERE id=?"
	row := db.DB.QueryRow(query, id)

	var article Article
	err := row.Scan(&article.ID, &article.Title, &article.Link)
	if err != nil {
		return Article{}, err
	}

	return article, nil
}
