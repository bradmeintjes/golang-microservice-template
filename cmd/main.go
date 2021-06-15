package main

import (
	"log"

	"github.com/caarlos0/env/v6"

	"sample-microservice-v2/internal/config"
	userService "sample-microservice-v2/internal/domain/user"
	"sample-microservice-v2/internal/repository/postgres"
	userStorage "sample-microservice-v2/internal/repository/postgres/user"
	"sample-microservice-v2/internal/transport/http"
	"sample-microservice-v2/internal/transport/http/user"
	userUsecase "sample-microservice-v2/internal/usecase/user"
)

func main() {
	c, err := conf()
	if err != nil {
		log.Fatalln(err)
	}

	p, err := postgres.NewConn(c.Postgres)
	if err != nil {
		log.Fatalln(err)
	}

	rUser := userStorage.NewRepository(p)
	sUser := userService.NewService(rUser, nil)
	uUser := userUsecase.NewUsecase(sUser)
	hUser := user.NewHandler(uUser)

	srv := http.NewServer(c.HTTP)
	srv.MountRoutes(hUser)
	srv.Listen()
}

func conf() (config.Config, error) {
	var c config.Config
	return c, env.Parse(&c)
}
