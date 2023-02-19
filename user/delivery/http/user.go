package http

import (
	"encoding/json"
	"github.com/event_bright/domain"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

type AuthHandler struct {
	AuthUseCase domain.AuthUseCase
}

func NewHTTPHandler(r *chi.Mux, authUseCase domain.AuthUseCase) {
	handler := &AuthHandler{
		AuthUseCase: authUseCase,
	}
	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/", handler.User)
	})
}

type ReqUser struct {
	domain.User
}

func (a *AuthHandler) User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := ReqUser{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
	}

	ctx := r.Context()
	user := domain.User(req.User)
	res := a.AuthUseCase.User(ctx, &user)

	er := json.NewEncoder(w).Encode(res)
	if er != nil {
		log.Println(er)
	}
}
