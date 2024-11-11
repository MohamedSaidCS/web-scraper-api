package main

import (
	"github.com/MohamedSaidCS/web-scraper-api/db"
	"github.com/MohamedSaidCS/web-scraper-api/routes"
	"github.com/MohamedSaidCS/web-scraper-api/scraper"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {
	db.InitDB()

	c := cron.New()

	c.AddFunc("@every 5m", func() {

		for _, site := range scraper.Sites {
			go scraper.ScrapeSite(site.URL, site.ArticleSelector, site.Extractor, 10)
		}
	})

	c.Start()

	server := gin.Default()

	routes.Init(server)

	server.Run(":8080")

}
