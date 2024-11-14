package scraper

import (
	"log"
	"net/http"

	"github.com/MohamedSaidCS/web-scraper-api/models"
	"github.com/PuerkitoBio/goquery"
)

type SiteInfo struct {
	URL             string
	ArticleSelector string
	Extractor       func(s *goquery.Selection) (string, string, string)
}

var SitesHTML = map[string]SiteInfo{
	"ArsTechnica": {
		URL:             "https://arstechnica.com/gadgets/",
		ArticleSelector: "article",
		Extractor: func(s *goquery.Selection) (string, string, string) {
			anchorTag := s.Find("h2 > a")

			title := anchorTag.Text()
			link, _ := anchorTag.Attr("href")
			timestamp, _ := s.Find("time").Attr("datetime")

			return title, link, timestamp
		},
	},
	"TechCrunch": {
		URL:             "https://techcrunch.com/latest/",
		ArticleSelector: ".post",
		Extractor: func(s *goquery.Selection) (string, string, string) {
			anchorTag := s.Find(".loop-card__title > a")

			title := anchorTag.Text()
			link, _ := anchorTag.Attr("href")
			timestamp, _ := s.Find("time").Attr("datetime")

			return title, link, timestamp
		},
	},
}

func ScrapeSiteHTML(website string, articleSelector string, extractor func(s *goquery.Selection) (string, string, string), maxArticles int) {
	res, err := http.Get(website)
	if err != nil || res.StatusCode != 200 {
		log.Println(err)
		return
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
		return
	}

	doc.Find(articleSelector).Each(func(i int, s *goquery.Selection) {
		if i > maxArticles-1 {
			return
		}

		title, link, timestamp := extractor(s)

		artice := models.Article{
			Title:    title,
			Link:     link,
			Timesamp: timestamp,
		}

		err = artice.Create()
		if err != nil {
			log.Println(err)
		}
	})
}
