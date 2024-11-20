package main

import (
	"github.com/MohamedSaidCS/web-scraper-api/db"
	"github.com/MohamedSaidCS/web-scraper-api/middlewares"
	"github.com/MohamedSaidCS/web-scraper-api/routes"
	"github.com/MohamedSaidCS/web-scraper-api/scraper"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/robfig/cron/v3"
)

func main() {
	db.InitDB()
	db.InitMongoDB()

	c := cron.New()

	c.AddFunc("@every 5m", func() {
		for _, site := range scraper.SitesRSS {
			go scraper.ScrapeSiteRSS(site, 10)
		}

		// for _, site := range scraper.SitesHTML {
		// 	go scraper.ScrapeSiteHTML(site.URL, site.ArticleSelector, site.Extractor, 10)
		// }
	})

	c.Start()

	server := gin.Default()

	server.Use(middlewares.RateLimiter())

	server.Use(middlewares.RequestLogger())

	routes.Init(server)

	server.Run(":8080")

}
