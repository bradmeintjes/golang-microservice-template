package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"sample-microservice-v2/internal/domain/user"
	userUsecase "sample-microservice-v2/internal/usecase/user"
)

type Handler struct {
	userUsecase userUsecase.Usecase
}

func NewHandler(userUsecase userUsecase.Usecase) Handler {
	return Handler{
		userUsecase: userUsecase,
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
		ID:   u.ID,
		Name: u.Name,
	}

	if err := h.userUsecase.Create(uD); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
