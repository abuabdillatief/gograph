package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/abuabdillatief/gograph/db"
	"github.com/abuabdillatief/gograph/graph/generated"
	models "github.com/abuabdillatief/gograph/graph/model"
)

//Database ...
var Database = db.Connect()

func (r *queryResolver) Videos(ctx context.Context) ([]*models.Video, error) {
	return Database.FindAll(), nil
}

func (r *queryResolver) Video(ctx context.Context, id string) (*models.Video, error) {
	return Database.FindByID(id), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
