package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"sample-microservice-v2/internal/config"
)

type Server struct {
	conf config.HTTP
	mux  chi.Router
}

func NewServer(c config.HTTP) Server {
	return Server{
		conf: c,
		mux:  chi.NewRouter(),
	}
}

type Handler interface {
	Router() chi.Router
}

func (s Server) MountRoutes(handlers ...Handler) {
	for _, h := range handlers {
		s.mux.Mount("/", h.Router())
	}
}

func (s Server) Listen() {
	addr := fmt.Sprintf("%s:%d", s.conf.Host, s.conf.Port)
	log.Fatalln(http.ListenAndServe(addr, s.mux))
}
