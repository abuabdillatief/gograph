package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/abuabdillatief/gqlgen-todos/db"
	"github.com/abuabdillatief/gqlgen-todos/graph/generated"
	"github.com/abuabdillatief/gqlgen-todos/graph/model"
)

var database = db.Connect()

func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	return database.Save(&input), nil

}
func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	return database.FindAll(), nil
}

func (r *queryResolver) Video(ctx context.Context, id string) (*model.Video, error) {
	return database.FindByID(id), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
