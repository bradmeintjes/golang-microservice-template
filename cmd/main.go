package main

import (
	"log"

	"webservice-template/internal/config"
	userService "webservice-template/internal/domain/user"
	"webservice-template/internal/repository/postgres"
	userStorage "webservice-template/internal/repository/postgres/user"
	"webservice-template/internal/transport/http"
	userHandler "webservice-template/internal/transport/http/user"
)

func main() {
	c, err := config.Parse()
	if err != nil {
		log.Fatalln(err)
	}

	p, err := postgres.NewConn(c.Postgres)
	if err != nil {
		log.Fatalln(err)
	}

	rUser := userStorage.NewRepository(p)
	sUser := userService.NewService(rUser, nil)
	hUser := userHandler.NewHandler(sUser)

	s := http.NewServer(c.HTTP)
	s.MountRoutes(hUser)

	if err = s.Serve(); err != nil {
		log.Fatalln(err)
	}
}
