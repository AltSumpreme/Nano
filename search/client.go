package search

import (
	"context"

	"github.com/AltSumpreme/Nano/search/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.SearchServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := pb.NewSearchServiceClient(conn)
	return &Client{conn, c}, nil

}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) PostBlog(ctx context.Context, name string, description string) (*Post, error) {
	r, err := c.service.PostBlog(
		ctx, &pb.PostBlogRequest{
			Name:        name,
			Description: description,
		},
	)

	if err != nil {
		return nil, err
	}

	return &Post{
		ID:          r.Post.Id,
		Name:        r.Post.Name,
		Description: r.Post.Description,
	}, nil

}
func (c *Client) GetPost(ctx context.Context, id string) (*Post, error) {
	r, err := c.service.GetPost(
		ctx, &pb.GetPostRequest{
			Id: id,
		},
	)
	if err != nil {
		return nil, err
	}
	return &Post{
		ID:          r.Post.Id,
		Name:        r.Post.Name,
		Description: r.Post.Description,
	}, nil

}

func (c *Client) GetPosts(ctx context.Context, ids []string, skip uint64, take uint64) ([]Post, error) {
	r, err := c.service.GetPosts(
		ctx, &pb.GetPostsRequest{
			Ids:  ids,
			Skip: skip,
			Take: take,
		},
	)

	if err != nil {
		return nil, err
	}
	posts := []Post{}
	for _, p := range r.Posts {
		posts = append(posts, Post{
			ID:          p.Id,
			Name:        p.Name,
			Description: p.Description,
		})
	}
	return posts, nil

}
