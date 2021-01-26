package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/abuabdillatief/gograph/common"
	"github.com/abuabdillatief/gograph/graph/generated"
	"github.com/abuabdillatief/gograph/graph/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	database := common.GetDB()
	var videos []*model.Video
	cursor, err := database.Videos.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(ctx) {
		var video *model.Video
		err = cursor.Decode(&video)
		if err != nil {
			log.Fatal(err)
		}
		//========================= +
		videos = append(videos, video)
	}
	return videos, nil

}

func (r *queryResolver) Video(ctx context.Context, id string) (*model.Video, error) {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	database := common.GetDB()
	var video *model.Video
	res := database.Videos.FindOne(ctx, bson.M{"_id": ObjectID})
	res.Decode(video)
	return video, nil

}

func (r *queryResolver) Audios(ctx context.Context) ([]*model.Audio, error) {
	database := common.GetDB()
	var audios []*model.Audio
	cursor, err := database.Audios.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var audio *model.Audio
		err = cursor.Decode(&audio)
		if err != nil {
			log.Fatal(err)
		}
		//========================= +
		audios = append(audios, audio)
	}
	return audios, nil

}

func (r *queryResolver) Audio(ctx context.Context, id string) (*model.Audio, error) {
	database := common.GetDB()
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	res := database.Audios.FindOne(ctx, bson.M{"_id": ObjectID})
	var audio *model.Audio
	res.Decode(&audio)
	return audio, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
