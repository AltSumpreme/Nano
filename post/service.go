package post

import (
	"context"
	"time"
)

type service interface {
	CreatePost(ctx context.Context, p Post) error
	GetPostByID(ctx context.Context, id string) (*Post, error)
	GetPostByName(ctx context.Context, name string) (*Post, error)
	GetPostByDescription(ctx context.Context, description string) (*Post, error)
	GetAllPosts(ctx context.Context) ([]*Post, error)
	DeletePost(ctx context.Context, id string) error
}

type Post struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	AccountID   string    `json:"account_id"`
	Likes       []Like    `json:"likes"`
	Comments    []Comment `json:"comments"`
}

type Like struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	PostID    string    `json:"post_id"`
	AccountID string    `json:"account_id"`
}

type Comment struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	PostID    string    `json:"post_id"`
	AccountID string    `json:"account_id"`
	Content   string    `json:"text"`
}
