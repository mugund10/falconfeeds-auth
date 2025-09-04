package main

import (
	"log"

	"github.com/mugund10/falconfeeds-auth/api"
	"github.com/mugund10/falconfeeds-auth/storage"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	// config init
	listenAddress := api.GetEnv("LIS_ADDR", "0.0.0.0:8080")
	dbAddress := api.GetEnv("DB_ADDR", "mongodb://localhost:27017")
	dbName := api.GetEnv("DB_NAME", "ff")

	// db init
	ops := options.Client().ApplyURI(dbAddress)
	client, err := mongo.Connect(ops)
	if err != nil {
		log.Panic("error connecting to db")
	}
	dbStore := storage.NewMongoUserStore(client.Database(dbName))

	// server init
	server := api.Newserver(listenAddress, dbStore)
	log.Printf("server running on http://%s", listenAddress)
	log.Fatal(server.Start())
}
