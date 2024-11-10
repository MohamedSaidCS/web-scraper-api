package main

import (
	"github.com/MohamedSaidCS/web-scraper-api/db"
	"github.com/MohamedSaidCS/web-scraper-api/routes"
	"github.com/MohamedSaidCS/web-scraper-api/scraper"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	for _, site := range scraper.Sites {
		scraper.ScrapeSite(site.URL, site.ArticleSelector, site.Extractor, 10)
	}

	server := gin.Default()

	routes.Init(server)

	server.Run(":8080")

}
