package lMongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateIndex(coll *mongo.Collection, keys bson.D) {
	index := mongo.IndexModel{
		Keys: keys,
	}
	_, err := coll.Indexes().CreateOne(context.TODO(), index)
	if err != nil {
		panic(err)
	}
}

func CreateSimpleIndex(coll *mongo.Collection, key string) {
	CreateIndex(coll, bson.D{
		{key, 1},
	})
}
