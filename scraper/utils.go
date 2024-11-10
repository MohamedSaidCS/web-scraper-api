package scraper

import "github.com/PuerkitoBio/goquery"

func ArsTechninaExtractor(s *goquery.Selection) (string, string, string) {
	anchorTag := s.Find("h2 > a")

	title := anchorTag.Text()
	link, _ := anchorTag.Attr("href")
	timestamp, _ := s.Find("time").Attr("datetime")

	return title, link, timestamp
}

func TechChrunchExtractor(s *goquery.Selection) (string, string, string) {
	anchorTag := s.Find(".loop-card__title > a")

	title := anchorTag.Text()
	link, _ := anchorTag.Attr("href")
	timestamp, _ := s.Find("time").Attr("datetime")

	return title, link, timestamp
}
