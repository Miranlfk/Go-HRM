package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InstanceMongo struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var Instance InstanceMongo

// ConnectDB connects to the MongoDB database using parameters from the .env file
func ConnectDB() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_DB_URI")
	dbName := os.Getenv("DB_NAME")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	db := client.Database(dbName)
	if err != nil {
		return err
	}

	Instance = InstanceMongo{
		Client: client,
		DB:     db,
	}
	return nil
}
