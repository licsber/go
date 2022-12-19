package lMongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"testing"
)

var client *mongo.Client

func Init() {
	uri := os.Getenv("L_MONGO_URI")
	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	client = c
}

func TestInsertOne(t *testing.T) {
	if client == nil {
		Init()
	}

	coll := client.Database("main").Collection("test")
	res, err := coll.InsertOne(context.TODO(), bson.M{
		"hello": "world",
	})
	fmt.Println(res)
	fmt.Println(err)
}

func TestFindOne(t *testing.T) {
	if client == nil {
		Init()
	}

	coll := client.Database("main").Collection("test")
	var res bson.M
	err := coll.FindOne(context.TODO(), bson.M{
		"hello": bson.M{
			"$exists": true,
		},
	}).Decode(&res)

	if err == mongo.ErrNoDocuments {
		fmt.Println("cannot find.")
		return
	}

	fmt.Println(res)
}

func TestUpdateOne(t *testing.T) {
	if client == nil {
		Init()
	}

	coll := client.Database("main").Collection("test")
	res, err := coll.UpdateOne(
		context.TODO(),
		bson.M{
			"hello": bson.M{
				"$exists": false,
			},
		},
		bson.M{
			"$set": bson.M{
				"update": true,
			},
		},
	)
	fmt.Println(err)
	fmt.Println(res.MatchedCount)
	fmt.Println(res.UpsertedCount)
	fmt.Println(res.ModifiedCount)
	fmt.Println(res.UpsertedID)
}

func TestUpsert(t *testing.T) {
	if client == nil {
		Init()
	}

	coll := client.Database("main").Collection("test")
	opts := options.Update().SetUpsert(true)
	res, err := coll.UpdateOne(
		context.TODO(),
		bson.M{
			"hello": bson.M{
				"$exists": false,
			},
		},
		bson.M{
			"$set": bson.M{
				"update": true,
			},
		},
		opts,
	)
	fmt.Println(err)
	fmt.Println(res.MatchedCount)
	fmt.Println(res.UpsertedCount)
	fmt.Println(res.ModifiedCount)
	fmt.Println(res.UpsertedID)
}
