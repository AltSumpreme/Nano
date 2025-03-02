package main

import "github.com/99designs/gqlgen/graphql"

type Server struct {
	accountClient *account.client
	postClient    *post.client
	searchClient  *search.client
}

func NewGraphQLServer(accountUrl, postUrl, searchUrl string) (*Server, error) {
	accountClient, err := account.NewClient(accountUrl)

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

func (s *Server) Account() QueryResolver {
	return &queryResolver{
		server: s,
	}
}

func (s *Server) ToexecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}
