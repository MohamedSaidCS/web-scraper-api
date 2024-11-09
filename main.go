package main

import (
	"github.com/MohamedSaidCS/web-scraper-api/db"
	"github.com/MohamedSaidCS/web-scraper-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.Init(server)

	server.Run(":8080")
}
