package main

import (
	"context"
	"log"

	"webservice-template/internal/config"
	userService "webservice-template/internal/domain/user"
	"webservice-template/internal/repository/postgres"
	userRepo "webservice-template/internal/repository/postgres/user"
	"webservice-template/internal/transport/http"
	userHandler "webservice-template/internal/transport/http/user"
)

func main() {

	c, err := config.Parse()
	if err != nil {
		log.Fatalln(err)
	}

	p, err := postgres.NewConn(context.Background(), c.Postgres)
	if err != nil {
		log.Fatalln(err)
	}
	defer p.Close()

	repo := userRepo.NewRepository(p)
	svc := userService.NewService(repo, nil)
	handler := userHandler.NewHandler(svc)

	s := http.NewServer(c.HTTP)
	s.MountRoutes(handler)

	if err = s.Serve(); err != nil {
		log.Fatalln(err)
	}
}
