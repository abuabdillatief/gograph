package graph

import "github.com/abuabdillatief/gqlgen-todos/graph/model"

// This file will not be regenerated automatically.
//go:generate go run github.com/99designs/gqlgen
// It serves as dependency injection for your app, add any dependencies you require here.

//Resolver ...
type Resolver struct {
	videos []*model.Video
}
