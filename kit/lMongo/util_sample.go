package lMongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func Sample(coll *mongo.Collection, num int, match bson.M, proj bson.D) (res []bson.M, err error) {
	if match == nil {
		match = bson.M{}
	}
	if proj == nil {
		proj = bson.D{}
	}

	cursor, err := coll.Aggregate(
		context.TODO(),
		bson.A{
			bson.M{"$match": match},
			bson.M{"$sample": bson.M{
				"size": num,
			}},
			bson.M{"$project": proj},
		},
	)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = cursor.Close(context.TODO())
		if err != nil {
			log.Fatalln(err)
		}
	}()

	for cursor.Next(context.TODO()) {
		var doc bson.M
		if err = cursor.Decode(&doc); err != nil {
			log.Fatal(err)
		}
		res = append(res, doc)
	}

	err = cursor.Err()
	return
}
