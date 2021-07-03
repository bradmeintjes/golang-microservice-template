package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
	"webservice-template/internal/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler interface {
	Router() chi.Router
}

type Server struct {
	conf config.HTTP
	mux  chi.Router
}

func NewServer(c config.HTTP) Server {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	return Server{
		conf: c,
		mux:  r,
	}
}

func (s Server) MountRoutes(handlers ...Handler) {
	for _, h := range handlers {
		s.mux.Mount("/", h.Router())
	}
}

func (s Server) Serve() error {
	addr := fmt.Sprintf("%s:%d", s.conf.Host, s.conf.Port)
	svr := &http.Server{
		Addr:         addr,
		Handler:      s.mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return gracefully(svr, func(s *http.Server) error {
		return s.ListenAndServe()
	})
}

// spawns a routine to wait for a interrupt signal and handle the shutdown gracefully
func gracefully(svr *http.Server, serve func(*http.Server) error) error {
	done := make(chan error)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		svr.SetKeepAlivesEnabled(false)
		if err := svr.Shutdown(ctx); err != nil {
			done <- fmt.Errorf("could not gracefully shutdown the server: %s", err)
		}
		done <- nil
	}()

	err := serve(svr)

	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("could not listen on %s: %v", svr.Addr, err)
	}

	return <-done
}
