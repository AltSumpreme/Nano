package main

import (
	"log"
	"time"

	auth "github.com/AltSumpreme/Nano/auth"
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

	var r auth.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = auth.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()

	log.Println("Listening on port 8080...")
	s := auth.NewService(r)
	log.Fatal(auth.ListenGrpc(s, 8080))
}
