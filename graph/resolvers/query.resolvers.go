package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/abuabdillatief/gograph/graph/generated"
	"github.com/abuabdillatief/gograph/graph/model"
)

func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	return DB.FindAll(), nil
}

func (r *queryResolver) Video(ctx context.Context, id string) (*model.Video, error) {
	return DB.FindByID(id), nil
}

func (r *queryResolver) Audios(ctx context.Context) ([]*model.Audio, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Audio(ctx context.Context, id string) (*model.Audio, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
