package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/DavidHODs/TechHUB-goGraph/auth"
	"github.com/DavidHODs/TechHUB-goGraph/graph/generated"
	"github.com/DavidHODs/TechHUB-goGraph/graph/model"
	database "github.com/DavidHODs/TechHUB-goGraph/postgres"
	"github.com/DavidHODs/TechHUB-goGraph/utils"
)

// returns created post author data
func (r *mutationResolver) CreatePost(ctx context.Context, input *model.NewPost) (*model.Post, error) {
	
	_ = ctx.Value("AuthToken")

	// authEmail, _ := auth.ParseToken(string(token))
	// if user == nil {
	// 	utils.HandleError(errors.New(("access denied")), false)
	// 	return &model.Post{}, errors.New(("access denied"))
	// }

	postAuthor := input.Author
	post := input.Body
	sharedPost := input.SharedBody
	postImage := input.SharedBody

	userID, userName, userEmail, err := database.ReturnUserDetails(postAuthor)
	if err != nil {
		utils.HandleError(err, false)
		return &model.Post{}, errors.New("something went wrong, try again later")
	}

	userDetails := model.User{
		ID:    userID,
		Name:  userName,
		Email: userEmail,
	}

	id, post, err := database.SavePost(postAuthor, post, sharedPost, postImage)
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
		Author:     &userDetails,
		SharedUser: nil,
		Likes:      nil,
		Dislikes:   nil,
		Tags:       nil,
	}, err
}

// returns created user data
func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	name := input.Name
	email := input.Email
	password := input.Password
	confirmPassword := input.Confirmpassword

	id, hashedP, err := database.SaveUser(name, email, password, confirmPassword)
	if err != nil {
		utils.HandleError(err, false)
	}

	token, _ := auth.GenerateToken(ctx, email)

	return &model.User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  string(hashedP),
		Token:     token,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// returns liked post data 
func (*mutationResolver) LikePost(ctx context.Context, input *model.UserPostID) (*model.Post, error) {
	uID := input.UserID
	pID := input.PostID

	userID, err := database.LikePostAndUpdateCount(uID, pID)
	if err != nil {
		utils.HandleError(err, false)
	}

	return &model.Post{
		ID:         pID,
		Body:       "",
		SharedBody: "",
		Image:      "",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		SharedAt:   time.Time{},
		Author:     &model.User{},
		SharedUser: []*model.UserID{},
		Likes:      []*model.UserID{
			{
				UserID: userID,
			},
		},
		Dislikes:   []*model.UserID{},
		Tags:       &model.Tag{},
	}, err
}

// returns minor details of user on succesful login
func (*mutationResolver) Login(ctx context.Context, input *model.LoginDetails) (*model.User, error) {
	email := input.Email
	password := input.Password 

	auth.GenerateToken(ctx, email)

	authenticated := auth.Authenticate(email, password)
	if !authenticated {
		utils.HandleError(errors.New("wrong email or password error"), false)
		return &model.User{}, errors.New("wrong email or password error") 
	}
	
	token, err := auth.GenerateToken(ctx, email)
	if err != nil{
		return &model.User{}, err
	}

	return &model.User{
		Email:     email,
		Token:     token,
	}, nil
}

// refreshes token of loggedin user 
func (r *mutationResolver) RefreshToken(ctx context.Context, input *model.Token) (*model.User, error) {
	tokenStr := input.Token

	email, err := auth.ParseToken(tokenStr)
	if err != nil {
		utils.HandleError(errors.New("access denied"), false)
		return &model.User{}, errors.New("access denied")
	}

	token, _ := auth.GenerateToken(ctx, email)
	
	return &model.User{
		Token:     token,
	}, nil
}

// returns unliked post data 
func (*mutationResolver) UnlikePost(ctx context.Context, input *model.UserPostID) (*model.Post, error) {
	panic("unimplemented")
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
