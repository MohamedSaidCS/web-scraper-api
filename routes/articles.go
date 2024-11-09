package routes

import (
	"github.com/MohamedSaidCS/web-scraper-api/models"
	"github.com/gin-gonic/gin"
)

func getArticles(context *gin.Context) {
	articles, err := models.GetArticles()
	if err != nil {
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(200, articles)
}
