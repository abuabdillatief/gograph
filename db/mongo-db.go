package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/abuabdillatief/gograph/graph/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

//Database ...
type Database struct {
	client *mongo.Client
}

const (
	//DATABASE ...
	DATABASE = "graphql"
)

//COLLECTION ...
var COLLECTION = map[string]string{
	"videos": "videos",
	"audios": "audios",
}

//Connect ...
func Connect() *Database {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	return &Database{
		client: client,
	}
}

//Save ...
func (db *Database) Save(video *model.NewVideo) *model.Video {
	collection := db.client.Database(DATABASE).Collection(COLLECTION["videos"])
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, video)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Video{
		ID:    res.InsertedID.(primitive.ObjectID).Hex(),
		Title: video.Title,
		URL:   video.URL,
	}
}

//FindAll ...
func (db *Database) FindAll() []*model.Video {
	var videos []*model.Video
	collection := db.client.Database(DATABASE).Collection(COLLECTION["videos"])
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Print(err)
	}
	for cursor.Next(ctx) {
		var video *model.Video
		err := cursor.Decode(&video)
		if err != nil {
			log.Fatal(err)
		}
		//==================== +
		videos = append(videos, video)
	}
	return videos
}

//FindByID ...
func (db *Database) FindByID(id string) *model.Video {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database(DATABASE).Collection(COLLECTION["videos"])
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	video := model.Video{}
	res.Decode(&video)
	return &video
}

//Delete ...
func (db *Database) Delete(id string) string {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database(DATABASE).Collection(COLLECTION["videos"])
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.DeleteOne(ctx, bson.M{"_id": ObjectID})
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("Deleted document: %v", res.DeletedCount)
}
