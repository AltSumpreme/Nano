package main

import (
	"context"
	"errors"
	"time"
)

type mutationResolver struct {
	server *Server
}

var (
	ErrInvalidParameter = errors.New("invalid parameter")
)

func (r *mutationResolver) createAccount(ctx context.Context, input AccountInput) (*Account, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	account, err := r.server.accountClient.PostAccount(ctx, input.Name)
	if err != nil {
		return nil, err
	}
	return &Account{
		ID:   account.ID,
		Name: account.Name,
	}, nil
}

func (r *mutationResolver) createPost(ctx context.Context, input PostInput) (*Post, error) {
	return r.server.postClient.createPost(ctx, input)
}
