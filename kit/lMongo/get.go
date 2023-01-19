package lMongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Get(cfg *Config) (client *mongo.Client) {
	if cfg.CTX == nil {
		cfg.CTX = context.TODO()
	}

	clientOptions := options.Client().ApplyURI(cfg.GetURI())
	client, err := mongo.Connect(cfg.CTX, clientOptions)
	if err != nil {
		panic(err)
	}

	return
}
