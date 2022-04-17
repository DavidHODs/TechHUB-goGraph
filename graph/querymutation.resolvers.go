package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/DavidHODs/TechHUB-goGraph/customerror"
	"github.com/DavidHODs/TechHUB-goGraph/graph/generated"
	"github.com/DavidHODs/TechHUB-goGraph/graph/model"
	database "github.com/DavidHODs/TechHUB-goGraph/postgres"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	var user database.User

	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password

	id, err := database.SaveUser()
	if err != nil {
		customerror.HandleError(err, false)
	}

	return &model.User{
		ID:        strconv.FormatInt(id, 10),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, err
}

func (r *queryResolver) Post(ctx context.Context) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
