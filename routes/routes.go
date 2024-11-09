package routes

import "github.com/gin-gonic/gin"

func Init(server *gin.Engine) {
	server.GET("/articles", getArticles)
}
