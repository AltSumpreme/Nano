package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AccountUrl string `envconfig:"ACCOUNT_SERVICE_URL"`
	PostUrl    string `envconfig:"POST_SERVICE_URL"`
	SearchUrl  string `envconfig:"SEARCH_SERVICE_URL"`
}

func main() {
	var cfg AppConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	s, err := NewGraphQLServer(cfg.AccountUrl, cfg.PostUrl, cfg.SearchUrl)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/graphql", handler.GraphQL(s.ToexecutableSchema()))
	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
