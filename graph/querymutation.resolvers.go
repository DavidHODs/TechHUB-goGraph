package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/DavidHODs/TechHUB-goGraph/graph/generated"
	"github.com/DavidHODs/TechHUB-goGraph/graph/model"
	database "github.com/DavidHODs/TechHUB-goGraph/postgres"
	"github.com/DavidHODs/TechHUB-goGraph/utils"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	post := input.Body
	sharedPost := input.SharedBody
	postImage := input.SharedBody

	id, post, err := database.SavePost(post, sharedPost, postImage)
	if err != nil {
		utils.HandleError(err, false)
	}

	return &model.Post{
		ID:         id,
		Body:       post,
		SharedBody: sharedPost,
		Image:      postImage,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		SharedAt:   time.Now(),
		Author:     nil,
		SharedUser: nil,
		Likes:      nil,
		Dislikes:   nil,
		Tags:       nil,
	}, err
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	name := input.Name
	email := input.Email
	password := input.Password
	confirmPassword := input.Confirmpassword

	id, hashedP, err := database.SaveUser(name, email, password, confirmPassword)
	if err != nil {
			utils.HandleError(err, false)
		}

	return &model.User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  string(hashedP),
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
