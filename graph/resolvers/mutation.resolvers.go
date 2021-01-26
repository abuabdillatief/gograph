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

func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	var database = common.GetDB()
	res, err := database.Videos.InsertOne(ctx, &input)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Video{
		ID:    res.InsertedID.(primitive.ObjectID).Hex(),
		Title: input.Title,
		URL:   input.URL,
	}, nil
}

func (r *mutationResolver) DeleteVideo(ctx context.Context, input string) (string, error) {
	ObjectID, err := primitive.ObjectIDFromHex(input)
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

func (r *mutationResolver) CreateAudio(ctx context.Context, input model.NewAudio) (*model.Audio, error) {
	var database = common.GetDB()
	res, err := database.Audios.InsertOne(ctx, &input)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Audio{
		ID:    res.InsertedID.(primitive.ObjectID).Hex(),
		Title: input.Title,
		URL:   input.URL,
	}, nil
}

func (r *mutationResolver) DeleteAudio(ctx context.Context, input string) (string, error) {
	ObjectID, err := primitive.ObjectIDFromHex(input)
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
