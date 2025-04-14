package main

import (
	"log"
	"time"

	"github.com/AltSumpreme/Nano/search"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)

	if err != nil {
		log.Fatal(err)
	}

	var r search.Repository
	retry.ForeverSleep(2*time.Second, func(int) (err error) {
		r, err = search.NewElasticRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()
	log.Println("Listening on port 8080...")
	s := search.NewService(r)
	log.Fatal(search.ListenGrpc(s, 8080))
}
