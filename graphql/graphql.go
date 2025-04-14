package main

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/AltSumpreme/Nano/auth"
	"github.com/AltSumpreme/Nano/search"
)

type Server struct {
	accountClient *auth.Client
	postClient    *post.Client
	searchClient  *search.Client
}

func NewGraphQLServer(accountUrl, postUrl, searchUrl string) (*Server, error) {
	accountClient, err := auth.NewClient(accountUrl)

	if err != nil {
		return nil, err
	}

	postClient, err := post.NewClient(postUrl)

	if err != nil {
		accountClient.Close()
		return nil, err
	}

	searchClient, err := search.NewClient(searchUrl)

	if err != nil {
		accountClient.Close()
		postClient.Close()
		return nil, err
	}

	return &Server{
		accountClient,
		postClient,
		searchClient,
	}, nil

}

func (s *Server) Mutation() MutationResolver {
	return &mutationResolver{
		server: s,
	}
}

func (s *Server) Query() QueryResolver {
	return &queryResolver{
		server: s,
	}
}

func (s *Server) Account() AccountResolver {
	return &queryResolver{
		server: s,
	}
}

func (s *Server) ToexecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
