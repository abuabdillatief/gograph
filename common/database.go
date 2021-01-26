package common

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

//DB ...
var DB *mongo.Collection

//ColsType ...
type ColsType struct {
	Audios *mongo.Collection
	Videos *mongo.Collection
}

//DBCollections ...
var DBCollections ColsType

//ConnectMongo ...
func ConnectMongo() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to mongodb")
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	DBCollections.Audios = client.Database("graphql").Collection("audios")
	DBCollections.Videos = client.Database("graphql").Collection("videos")

	_, err = DBCollections.Audios.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bsonx.Doc{{Key: "audio", Value: bsonx.Int64(1)}},
			Options: options.Index().SetUnique(false),
		},
	)

	_, err = DBCollections.Videos.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bsonx.Doc{{Key: "video", Value: bsonx.Int64(1)}},
			Options: options.Index().SetUnique(false),
		},
	)
}

//GetDB ...
func GetDB() ColsType {
	return DBCollections
}
