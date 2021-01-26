package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/abuabdillatief/gqlgen-todos/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

//VideoDB ...
type VideoDB interface {
	Save(video *model.Video)
	FindAll() []*model.Video
}

type database struct {
	client *mongo.Client
}

const (
	//DATABASE ...
	DATABASE = "graphql"

	//COLLECTION ...
	COLLECTION = "videos"
)

//New ...
func New() VideoDB {
	//mongodb+srv://USERNAME:PASSWORD@HOST:PORT
	// MONGODB := os.Getenv("MONGODB")
	clientOptions := options.Client().ApplyURI("mongodb+srv://root:@localhost:27017")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	mongo.Connect(ctx, clientOptions)
	dbClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return &database{
		client: dbClient,
	}
}

//Save ...
func (db *database) Save(video *model.Video) {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), video)
	if err != nil {
		log.Fatal(err)
	}
}

//FindAll ...
func (db *database) FindAll() []*model.Video {
	var result []*model.Video
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var video *model.Video
		err := cursor.Decode(&video)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, video)
	}
	return result

}
