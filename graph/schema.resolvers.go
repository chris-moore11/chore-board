package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cmoore/chore-board/graph/generated"
	"cmoore/chore-board/graph/model"
	"context"
	"fmt"
	"math/rand"
)

func (r *choreResolver) User(ctx context.Context, obj *model.Chore) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

func (r *mutationResolver) CreateChore(ctx context.Context, input model.NewChore) (*model.Chore, error) {
	chore := &model.Chore{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.chores = append(r.chores, chore)
	return chore, nil
}

func (r *queryResolver) Chores(ctx context.Context) ([]*model.Chore, error) {
	return r.chores, nil
}

// Chore returns generated.ChoreResolver implementation.
func (r *Resolver) Chore() generated.ChoreResolver { return &choreResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type choreResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
