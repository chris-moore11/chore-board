package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cmoore/chore-board/db"
	"cmoore/chore-board/graph/generated"
	"cmoore/chore-board/graph/model"
	"context"
)

func (r *mutationResolver) RotateForward(ctx context.Context) (bool, error) {
	var (
		id      int
		choreId int
	)
	users, err := db.GlobalInstance.Query("SELECT id, choreId FROM users")
	if err != nil {
		panic(err)
	}
	defer users.Close()

	var prevRowId int = -1
	var firstChoreId = -1
	updateStatement := "UPDATE users SET choreId=$1 WHERE id=$2;"

	for users.Next() {
		err := users.Scan(&id, &choreId)
		if err != nil {
			panic(err)
		}
		if prevRowId > 0 {
			db.GlobalInstance.Exec(updateStatement, choreId, prevRowId)
		} else {
			firstChoreId = choreId
		}
		prevRowId = id
	}
	if prevRowId > 0 {
		db.GlobalInstance.Exec(updateStatement, firstChoreId, prevRowId)
	}

	return true, nil
}

func (r *mutationResolver) RotateBackward(ctx context.Context) (bool, error) {
	var (
		id      int
		choreId int
	)
	users, err := db.GlobalInstance.Query("SELECT id, choreId FROM users")
	if err != nil {
		panic(err)
	}
	defer users.Close()

	var prevChoreId int = -1
	var firstRowId = -1
	updateStatement := "UPDATE users SET choreId=$1 WHERE id=$2;"

	for users.Next() {
		err := users.Scan(&id, &choreId)
		if err != nil {
			panic(err)
		}
		if prevChoreId > 0 {
			db.GlobalInstance.Exec(updateStatement, prevChoreId, id)
		} else {
			firstRowId = id
		}
		prevChoreId = choreId
	}
	if prevChoreId > 0 {
		db.GlobalInstance.Exec(updateStatement, prevChoreId, firstRowId)
	}

	return true, nil
}

func (r *queryResolver) Chores(ctx context.Context) ([]*model.Chore, error) {
	var (
		id          int
		text        string
		done        bool
		description string
		image       string
	)
	rows, err := db.GlobalInstance.Query("SELECT * FROM chores")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var chores []*model.Chore

	for rows.Next() {
		err := rows.Scan(&id, &text, &done, &description, &image)
		if err != nil {
			panic(err)
		}
		chore := &model.Chore{ID: id, Text: text, Done: done, Description: description, Image: image}
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
		id       int
		name     string
		email    string
		demerits int
		choreId  int
		admin    bool
	)
	rows, err := db.GlobalInstance.Query("SELECT * FROM users ORDER BY choreId")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []*model.User

	for rows.Next() {
		err := rows.Scan(&id, &name, &email, &demerits, &choreId, &admin)
		if err != nil {
			panic(err)
		}
		user := &model.User{ID: id, Name: name, Email: email, Demerits: demerits, ChoreID: choreId, Admin: admin}
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
const MAX_CHORE_ID = 8
