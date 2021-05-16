package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/medwing/insights-service/graph/generated"
	"github.com/medwing/insights-service/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := model.User{
		Name: input.Name,
	}
	r.DB.Create(&user)
	return &user, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		UserID: input.UserID,
	}
	result := r.DB.Create(&todo)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return todo, nil
	}
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.EditTodo) (*model.Todo, error) {
	todo := model.Todo{ID: input.ID}
	result := r.DB.First(&todo).Update("text", input.Text)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.DeleteTodo) (*model.Todo, error) {
	result := r.DB.Delete(&model.Todo{}, input.ID)
	if result.Error != nil {
		return nil, result.Error
	} else {
		deletedTodo := &model.Todo{
			ID: input.ID,
		}
		return deletedTodo, nil
	}
}

func (r *queryResolver) Todo(ctx context.Context, input model.FetchTodo) (*model.Todo, error) {
	todo := model.Todo{}
	result := r.DB.First(&todo, input.ID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	result := r.DB.Find(&r.todos)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.todos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
