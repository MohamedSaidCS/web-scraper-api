package routes

import (
	"github.com/MohamedSaidCS/web-scraper-api/models"
	"github.com/gin-gonic/gin"
)

func getArticles(context *gin.Context) {
	pageParam := context.Query("page")
	perPageParam := context.Query("per_page")
	articles, page, perPage, pages, total, err := models.GetArticles(pageParam, perPageParam)
	if err != nil {
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(200, gin.H{
		"data": articles,
		"paging": gin.H{
			"page":     page,
			"per_page": perPage,
			"pages":    pages,
			"total":    total,
		},
	})

}
