package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type Resource struct {
	Db *mongo.Database
}

func InitResource() (*Resource, error) {
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	dbName := os.Getenv("MONGO_DBNAME")
	dbUser := os.Getenv("MONGO_USER")
	dbPassword := os.Getenv("MONGO_PASSWORD")

	credential := options.Credential{
		Username: dbUser,
		Password: dbPassword,
	}

	uri := fmt.Sprintf("mongodb://%s:%s", host, port)
	clientOptions := options.Client()
	clientOptions.ApplyURI(uri).SetAuth(credential)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	return &Resource{
		Db: client.Database(dbName),
	}, nil
}
