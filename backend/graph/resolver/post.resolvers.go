package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"backend/graph"
	"backend/graph/model"
	"context"
	"fmt"

	"github.com/google/uuid"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.CreatePostInput) (*model.CreatePostPayload, error) {
	var user model.User
	if err := r.DB.First(&user, input.AuthorID).Error; err != nil {
		return &model.CreatePostPayload{
			Post: &model.Post{},
			Error: &model.Error{
				Message: err.Error(),
			},
		}, err
	}
	id := uuid.New().String()
	post := model.Post{
		ID:      id,
		Content: input.Content,
		Author:  &user,
	}

	if err := r.DB.Create(&post).Error; err == nil {
		return &model.CreatePostPayload{
			Post: &post,
			Error: &model.Error{
				Message: err.Error(),
			},
		}, err
	}
	return &model.CreatePostPayload{
		Post: &post,
	}, nil
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, input model.UpdatePostInput) (*model.UpdatePostPaylod, error) {
	panic(fmt.Errorf("not implemented: UpdatePost - updatePost"))
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, input model.ModelInputID) (*model.DeletePostPayload, error) {
	panic(fmt.Errorf("not implemented: DeletePost - deletePost"))
}

// GetPost is the resolver for the getPost field.
func (r *queryResolver) GetPost(ctx context.Context, input model.ModelInputID) (*model.GetPostPayload, error) {
	var post model.Post
	err := r.DB.First(&post, input.ID).Error
	return &model.GetPostPayload{
		Post: &post,
	}, err
}

// GetPosts is the resolver for the getPosts field.
func (r *queryResolver) GetPosts(ctx context.Context, input model.ConnectionInput) (*model.GetPostsPayload, error) {
	var posts []*model.Post
	var totalCount int64
	err := r.DB.Find(posts).Count(&totalCount).Error
	return &model.GetPostsPayload{
		Posts: posts,
		PageInfo: &model.CommonPageInfo{
			TotalCount:      int(totalCount),
			HasNextPage:     false,
			HasPreviousPage: false,
		},
	}, err
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
