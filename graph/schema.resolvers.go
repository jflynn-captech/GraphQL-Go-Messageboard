package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"fmt"

	"messageboard.example.graphql/.gen/messageboardDB/public/model"
	gqlmodel "messageboard.example.graphql/graph/model"
	"messageboard.example.graphql/internal/dataloaders"
)

// AuthorUser is the resolver for the authorUser field.
func (r *commentResolver) AuthorUser(ctx context.Context, obj *gqlmodel.Comment) (*gqlmodel.User, error) {
	return dataloaders.LoadUser(ctx, fmt.Sprint(obj.AuthorUserID))
}

// AddPost is the resolver for the addPost field.
func (r *mutationResolver) AddPost(ctx context.Context, add gqlmodel.AddNewPostInput) (*gqlmodel.Post, error) {
	post, err := r.PostService.AddPost(&model.Post{
		Text:          &add.Text,
		AuthorUsersID: int32(r.LoggedInUserId),
	})

	if err != nil {
		return nil, err
	}

	return &gqlmodel.Post{
		ID:           fmt.Sprint(post.ID),
		Text:         *post.Text,
		AuthorUserID: fmt.Sprint(post.AuthorUsersID),
	}, nil
}

// AddComment is the resolver for the addComment field.
func (r *mutationResolver) AddComment(ctx context.Context, add gqlmodel.AddNewCommentInput) (*gqlmodel.Comment, error) {
	post, err := r.PostService.GetPostById(add.PostID)

	if err != nil {
		return nil, err
	}

	comment, err := r.PostService.AddComment(&model.Comment{
		PostID:        post.ID,
		Text:          &add.Text,
		AuthorUsersID: int32(r.LoggedInUserId),
	})

	if err != nil {
		return nil, err
	}

	return &gqlmodel.Comment{
		ID:           fmt.Sprint(comment.ID),
		Text:         *comment.Text,
		AuthorUserID: int(comment.AuthorUsersID),
	}, err
}

// Note this is not a root resolver, it is a resolver
func (r *postResolver) AuthorUser(ctx context.Context, obj *gqlmodel.Post) (*gqlmodel.User, error) {
	return dataloaders.LoadUser(ctx, obj.AuthorUserID)
}

// Resolver for Post Comments
func (r *postResolver) Comments(ctx context.Context, obj *gqlmodel.Post, limit int) ([]*gqlmodel.Comment, error) {
	return dataloaders.LoadPostComment(ctx, obj.ID, limit)
}

// GetUsers is the resolver for the getUsers field.
func (r *queryResolver) GetUsers(ctx context.Context) ([]*gqlmodel.User, error) {
	dmUsers, err := r.UserService.GetUsers()

	if err != nil {
		return nil, err
	}

	var modelUsers []*gqlmodel.User
	for _, dmUser := range dmUsers {
		modelUsers = append(modelUsers, &gqlmodel.User{
			ID:   fmt.Sprint(dmUser.ID),
			Name: fmt.Sprint(dmUser.Name),
		})
	}

	return modelUsers, nil
}

// GetPosts is the resolver for the getPosts field.
func (r *queryResolver) GetPosts(ctx context.Context) ([]*gqlmodel.Post, error) {
	dmPosts, err := r.PostService.GetPosts()

	if err != nil {
		return nil, err
	}

	var modelPosts []*gqlmodel.Post

	for _, dmPost := range dmPosts {

		post := gqlmodel.Post{
			ID:           fmt.Sprint(dmPost.ID),
			AuthorUserID: fmt.Sprint(dmPost.AuthorUsersID),
			Text:         *dmPost.Text,
		}

		modelPosts = append(modelPosts, &post)
	}
	return modelPosts, nil
}

// Comment returns CommentResolver implementation.
func (r *Resolver) Comment() CommentResolver { return &commentResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Post returns PostResolver implementation.
func (r *Resolver) Post() PostResolver { return &postResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type commentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
