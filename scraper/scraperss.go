package scraper

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/MohamedSaidCS/web-scraper-api/models"
)

var SitesRSS = []string{
	"https://feeds.arstechnica.com/arstechnica/index",
	"https://techcrunch.com/feed/",
}

type SiteFeed struct {
	Items []Item `xml:"channel>item"`
}

type Item struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	PubDate string `xml:"pubDate"`
}

func (i Item) Timestamp() string {
	t, _ := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", i.PubDate)
	return t.Format(time.RFC3339)
}

func ScrapeSiteRSS(website string, maxArticles int) {
	res, err := http.Get(website)
	if err != nil || res.StatusCode != 200 {
		log.Println(err)
		return
	}

	defer res.Body.Close()

	xmlBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var feed SiteFeed

	xml.Unmarshal(xmlBytes, &feed)

	for i, item := range feed.Items {
		if i > maxArticles-1 {
			return
		}

		article := models.Article{
			Title:    item.Title,
			Link:     item.Link,
			Timesamp: item.Timestamp(),
		}

		err = article.Create()
		if err != nil {
			log.Println(err)
		}
	}

}
