package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cmoore/chore-board/db"
	"cmoore/chore-board/graph/generated"
	"cmoore/chore-board/graph/model"
	"context"
)

const MAX_CHORE_ID = 8

func (r *mutationResolver) RotateForward(ctx context.Context) (bool, error) {
	var (
		id      int
		choreId int
	)
	rows, err := db.GlobalInstance.Query("SELECT id, choreId FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &choreId)
		if err != nil {
			panic(err)
		}
		choreId++
		if choreId > MAX_CHORE_ID {
			choreId = 1
		}
		updateStatement := "UPDATE users SET choreId=$1 WHERE id=$2;"
		db.GlobalInstance.Exec(updateStatement, choreId, id)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return true, nil
}

func (r *mutationResolver) RotateBackward(ctx context.Context) (bool, error) {
	var (
		id      int
		choreId int
	)
	rows, err := db.GlobalInstance.Query("SELECT id, choreId FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &choreId)
		if err != nil {
			panic(err)
		}
		choreId--
		if choreId < 1 {
			choreId = 8
		}
		updateStatement := "UPDATE users SET choreId=$1 WHERE id=$2;"
		db.GlobalInstance.Exec(updateStatement, choreId, id)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return true, nil
}

func (r *queryResolver) Chores(ctx context.Context) ([]*model.Chore, error) {
	var (
		id       int
		text     string
		done     bool
		image    string
		tutorial string
	)
	rows, err := db.GlobalInstance.Query("SELECT * FROM chores")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var chores []*model.Chore

	for rows.Next() {
		err := rows.Scan(&id, &text, &done, &image, &tutorial)
		if err != nil {
			panic(err)
		}
		chore := &model.Chore{ID: id, Text: text, Done: done, Image: image, Tutorial: tutorial}
		chores = append(chores, chore)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return chores, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var (
		id      int
		name    string
		email   string
		image   string
		choreId int
		admin   bool
	)
	rows, err := db.GlobalInstance.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []*model.User

	for rows.Next() {
		err := rows.Scan(&id, &name, &email, &image, &choreId, &admin)
		if err != nil {
			panic(err)
		}
		user := &model.User{ID: id, Name: name, Email: email, Image: image, ChoreID: choreId, Admin: admin}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
