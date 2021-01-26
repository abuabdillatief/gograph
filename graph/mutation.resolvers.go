package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/abuabdillatief/gograph/graph/generated"
	models "github.com/abuabdillatief/gograph/graph/model"
)

func (r *mutationResolver) CreateVideo(ctx context.Context, input models.NewVideo) (*models.Video, error) {
	return Database.Save(&input), nil
}

func (r *mutationResolver) DeleteVideo(ctx context.Context, input string) (string, error) {
	return Database.Delete(input), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
