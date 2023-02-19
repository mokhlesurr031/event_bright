package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/event_bright/domain"
	"github.com/event_bright/domain/dto"
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
		r.Post("/login", handler.SignIn)
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

type ReqSignIn struct {
	dto.SignIn
}

func (a *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	req := ReqSignIn{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
	}

	ctx := r.Context()
	signIn := dto.SignIn(req.SignIn)
	resp, err := a.AuthUseCase.SignIn(ctx, &signIn)
	if err != nil {
		log.Println(err)
	}
	er := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(er)
	}
}
