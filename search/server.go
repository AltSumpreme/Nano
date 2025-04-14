package search

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/AltSumpreme/Nano/search/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	pb.UnimplementedSearchServiceServer
	service Service
}

func ListenGrpc(service Service, port int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server, &grpcServer{service: service})
	reflection.Register(server)
	return server.Serve(listener)
}

func (s *grpcServer) PostBlog(ctx context.Context, r *pb.PostBlogRequest) (*pb.PostBlogResponse, error) {
	post, err := s.service.PostBlog(ctx, r.Name, r.Description)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.PostBlogResponse{Post: &pb.Post{
		Id:          post.ID,
		Name:        post.Name,
		Description: post.Description,
	}}, nil

}

func (s *grpcServer) GetPost(ctx context.Context, r *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	post, err := s.service.GetPost(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetPostResponse{Post: &pb.Post{
		Id:          post.ID,
		Name:        post.Name,
		Description: post.Description,
	}}, nil
}

func (s *grpcServer) GetPosts(ctx context.Context, r *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	var res []Post
	var err error
	if r.Query != "" {
		res, err = s.service.SearchPosts(ctx, r.Query, r.Skip, r.Take)
	} else if len(r.Ids) != 0 {
		res, err = s.service.GetPostsByID(ctx, r.Ids)
	} else {
		res, err = s.service.GetPosts(ctx, r.Skip, r.Take)
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	posts := []*pb.Post{}
	for _, p := range res {
		posts = append(posts, &pb.Post{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
		})
	}
	return &pb.GetPostsResponse{Posts: posts}, nil

}
