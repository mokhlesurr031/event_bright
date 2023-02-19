package http

import (
	"context"
	"encoding/json"
	"github.com/event_bright/domain"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strings"
)

type EventHandler struct {
	EventUseCase domain.EventUseCase
}

func NewHTTPHandler(r *chi.Mux, eventUseCase domain.EventUseCase) {
	handler := &EventHandler{
		EventUseCase: eventUseCase,
	}
	r.Route("/api/v1/event", func(r chi.Router) {
		r.Post("/", handler.Event)
	})
}

type ReqEvent struct {
	domain.Event
}

func (e *EventHandler) Event(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check that the request includes a valid JWT token
	tokenString := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	if tokenString == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	ctx := context.WithValue(r.Context(), "token", tokenString)

	req := ReqEvent{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
	}
	event := domain.Event(req.Event)
	res, err := e.EventUseCase.Event(ctx, &event)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
