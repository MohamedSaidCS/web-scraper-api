package utils

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/MohamedSaidCS/web-scraper-api/db"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoWriter struct {
	dbName string
}

func (mw *MongoWriter) Write(p []byte) (n int, err error) {
	_, err = db.Client.Database(mw.dbName).Collection("errorlogs").InsertOne(context.TODO(), bson.D{
		{Key: "error", Value: string(p)},
		{Key: "timestamp", Value: time.Now()},
	})
	if err != nil {
		return 0, err
	}

	return len(p), nil
}

func InitErrorLogger() {
	multiwriter := io.MultiWriter(os.Stdout, &MongoWriter{dbName: os.Getenv("MONGO_DB")})
	log.SetOutput(multiwriter)
}
