package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"webservice-template/internal/domain/user"
)

type Handler struct {
	userSvc user.Service
}

func NewHandler(userSvc user.Service) Handler {
	return Handler{
		userSvc: userSvc,
	}
}

func (h Handler) Router() chi.Router {
	r := chi.NewRouter()
	r.Post("/user", h.Create)
	return r
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uD := user.User{
		Name: u.Name,
	}

	if err := h.userSvc.Create(uD); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
