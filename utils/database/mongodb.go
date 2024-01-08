package database

import (
	"context"
	"fmt"
	"gomongo/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB(c config.MongoConfig) (*mongo.Database, error) {
	mongoURL := fmt.Sprintf("mongodb://%s",
		c.DBLINK,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal("failed to connect", err.Error())
		return nil, err
	}

	database := client.Database(c.DBNAME)

	return database, nil

}
