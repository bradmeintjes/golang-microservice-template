package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"webservice-template/internal/config"

	"github.com/go-chi/chi/v5"
)

type Handler interface {
	Router() chi.Router
}

type key int

const (
	requestIDKey key = 0
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

func (s Server) MountRoutes(handlers ...Handler) {
	for _, h := range handlers {
		s.mux.Mount("/", h.Router())
	}
}

func (s Server) Serve() error {
	addr := fmt.Sprintf("%s:%d", s.conf.Host, s.conf.Port)
	log.Println("starting server at " + addr)
	svr := &http.Server{
		Addr:         addr,
		Handler:      tracing(nextRequestID)(logging()(s.mux)),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return gracefully(svr, func(s *http.Server) error {
		return s.ListenAndServe()
	})
}

func nextRequestID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func logging() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				requestID, ok := r.Context().Value(requestIDKey).(string)
				if !ok {
					requestID = "unknown"
				}
				log.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
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
