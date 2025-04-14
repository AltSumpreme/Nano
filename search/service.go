package search

import (
	"context"

	"github.com/segmentio/ksuid"
)

type Service interface {
	PostBlog(ctx context.Context, name string, description string) (*Post, error)
	GetPost(ctx context.Context, id string) (*Post, error)
	GetPosts(ctx context.Context, skip uint64, take uint64) ([]Post, error)
	GetPostsByID(ctx context.Context, ids []string) ([]Post, error)
	SearchPosts(ctx context.Context, query string, skip uint64, take uint64) ([]Post, error)
}

type Post struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type searchService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &searchService{repository: r}
}

func (s *searchService) PostBlog(ctx context.Context, name string, description string) (*Post, error) {
	p := &Post{
		ID:          ksuid.New().String(),
		Name:        name,
		Description: description,
	}
	if err := s.repository.PutPost(ctx, *p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *searchService) GetPost(ctx context.Context, id string) (*Post, error) {
	res, err := s.repository.GetPostByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *searchService) GetPosts(ctx context.Context, skip uint64, take uint64) ([]Post, error) {
	if take > 50 || (skip == 0 && take == 0) {
		take = 50
	}
	res, err := s.repository.ListPosts(ctx, skip, take)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (s *searchService) GetPostsByID(ctx context.Context, ids []string) ([]Post, error) {
	res, err := s.repository.ListPostsWithIDS(ctx, ids)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *searchService) SearchPosts(ctx context.Context, query string, skip uint64, take uint64) ([]Post, error) {
	if take > 50 || (skip == 0 && take == 0) {
		take = 50
	}
	res, err := s.repository.SearchPosts(ctx, query, skip, take)
	if err != nil {
		return nil, err
	}
	return res, nil
}
