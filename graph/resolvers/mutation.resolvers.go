package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/abuabdillatief/gograph/graph/generated"
	"github.com/abuabdillatief/gograph/graph/model"
)

func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	return DB.Save(&input), nil
}

func (r *mutationResolver) DeleteVideo(ctx context.Context, input string) (string, error) {
	return DB.Delete(input), nil
}

func (r *mutationResolver) CreateAudio(ctx context.Context, input model.NewAudio) (*model.Audio, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAudio(ctx context.Context, input string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
