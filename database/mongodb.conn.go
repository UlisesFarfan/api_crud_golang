package database

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	database = "golang_test"
)

// Database connection
func GetColllection(collection string) *mongo.Collection {
	godotenv.Load()
	// Change url for tests
	url := os.Getenv("MONGO_URL")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}
	return client.Database(database).Collection(collection)
}
