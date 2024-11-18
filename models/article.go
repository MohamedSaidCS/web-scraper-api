package models

import (
	"fmt"
	"math"
	"strconv"

	"github.com/MohamedSaidCS/web-scraper-api/db"
	"github.com/lib/pq"
)

type Article struct {
	ID       int64  `json:"id"`
	Title    string `json:"title" binding:"required"`
	Link     string `json:"link" binding:"required"`
	Timesamp string `json:"timestamp"`
}

func (a *Article) Create() error {
	query := "INSERT INTO articles (title, link, timestamp) VALUES ($1, $2, $3) RETURNING id, timestamp"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRow(a.Title, a.Link, a.Timesamp).Scan(&a.ID, &a.Timesamp)
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code == "23505" {
			return fmt.Errorf("article with link %s already exists", a.Link)
		}
		return err
	}

	return nil
}

func GetArticles(page string, perPage string) ([]Article, int, int, int, int, error) {
	pageInt, _ := strconv.Atoi(page)
	perPageInt, err := strconv.Atoi(perPage)
	if err != nil || perPageInt <= 0 {
		perPageInt = 5
	}
	pages := 1
	total := -1

	var query string
	if pageInt >= 1 {
		query = fmt.Sprintf("SELECT *, COUNT (*) OVER() FROM articles LIMIT %d OFFSET %d", perPageInt, (pageInt-1)*perPageInt)
	} else {
		query = "SELECT * FROM articles"
	}

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	defer rows.Close()

	articles := []Article{}

	for rows.Next() {
		var article Article
		if pageInt >= 1 {
			err = rows.Scan(&article.ID, &article.Title, &article.Link, &article.Timesamp, &total)
		} else {
			err = rows.Scan(&article.ID, &article.Title, &article.Link, &article.Timesamp)
		}
		if err != nil {
			return nil, 0, 0, 0, 0, err
		}

		articles = append(articles, article)
	}

	if total == -1 {
		total = len(articles)
		perPageInt = total
		if total == 0 {
			pageInt = 0
			pages = 0
		} else {
			pageInt = 1
		}
	} else {
		pages = int(math.Ceil(float64(total) / float64(perPageInt)))
	}

	return articles, pageInt, perPageInt, pages, total, nil
}
