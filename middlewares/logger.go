package middlewares

import (
	"context"
	"os"
	"time"

	"github.com/MohamedSaidCS/web-scraper-api/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var dBName = os.Getenv("MONGO_DB")

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		method := c.Request.Method
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		db.Client.Database(dBName).Collection("requestlogs").InsertOne(context.TODO(), bson.D{
			{Key: "method", Value: method},
			{Key: "path", Value: path},
			{Key: "status_code", Value: statusCode},
			{Key: "client_ip", Value: clientIP},
			{Key: "timestamp", Value: time.Now()},
		})

	}
}
