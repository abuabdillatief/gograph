package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/abuabdillatief/gograph/common"
	"github.com/abuabdillatief/gograph/graph/generated"
	"github.com/abuabdillatief/gograph/graph/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func (r *mutationResolver) CreateVideo(ctx context.Context, newVideo model.NewVideo) (*model.Video, error) {
	var database = common.GetDB()
	user := model.User{
		ID:   newVideo.UserID,
		Name: fmt.Sprintf("user_%v", newVideo.UserID),
	}
	res, err := database.Users.InsertOne(ctx, &user)
	if err != nil {
		log.Fatal(err)
	}
	res, err = database.Videos.InsertOne(ctx, &newVideo)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Video{
		ID:     res.InsertedID.(primitive.ObjectID).Hex(),
		Title:  newVideo.Title,
		URL:    newVideo.URL,
		Author: &user,
	}, nil
}

func (r *mutationResolver) DeleteVideo(ctx context.Context, id string) (string, error) {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	var database = common.GetDB()
	res, err := database.Videos.DeleteOne(ctx, bson.M{"_id": ObjectID})
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("Deleted document: %v", res.DeletedCount), nil
}

func (r *mutationResolver) CreateAudio(ctx context.Context, newAudio model.NewAudio) (*model.Audio, error) {
	var database = common.GetDB()
	user := model.User{
		ID:   newAudio.UserID,
		Name: fmt.Sprintf("user_%v", newAudio.UserID),
	}
	res, err := database.Audios.InsertOne(ctx, &newAudio)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Audio{
		ID:     res.InsertedID.(primitive.ObjectID).Hex(),
		Title:  newAudio.Title,
		URL:    newAudio.URL,
		Author: &user,
	}, nil
}

func (r *mutationResolver) DeleteAudio(ctx context.Context, id string) (string, error) {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	var database = common.GetDB()
	res, err := database.Audios.DeleteOne(ctx, bson.M{"_id": ObjectID})
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("Deleted document: %v", res.DeletedCount), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
