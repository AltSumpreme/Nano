package main

import "context"

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) createAccount(ctx context.Context, in AccountInput) (*Account, error) {
	return r.server.accountClient.createAccount(ctx, in)
}

func (r *mutationResolver) createPost(ctx context.Context, in PostInput) (*Post, error) {
	return r.server.postClient.createPost(ctx, in)
}
