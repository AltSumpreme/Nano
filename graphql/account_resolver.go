package main

import "context"

type accountResolver struct {
	server *Server
}

func (r *accountResolver) Post(ctx context.Context, pagiation *PaginationInput, id *string) ([]*Post, error) {
	return r.server.accountClient.GetAccount(ctx)
}
