package main

import "context"

type queryResolver struct {
	server *Server
}

func (r *queryResolver) Account(ctx context.Context, pagiation *PaginationInput, id *string) ([]*Account, error) {
	return r.server.accountClient.GetAccount(ctx)
}

func (r *queryResolver) Post(ctx context.Context, pagination *PaginationInput, id *string) ([]*Post, error) {
	return r.server.postClient.GetPost(ctx)
}
