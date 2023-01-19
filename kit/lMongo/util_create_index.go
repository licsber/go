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

func CreateSimpleIndex(coll *mongo.Collection, keys []string) {
	if len(keys) == 0 {
		panic("try to create empty index.")
	}

	var d bson.D
	for _, key := range keys {
		d = append(d, bson.E{Key: key, Value: 1})
	}

	CreateIndex(coll, d)
}
