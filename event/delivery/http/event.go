package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"

	"github.com/event_bright/domain"
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
		r.Get("/list", handler.EventList)
		r.Get("/list/me", handler.MyEventList)
	})
}

type ReqEvent struct {
	domain.Event
}

var (
	TokenKey = "token"
)

func (e *EventHandler) Event(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check that the request includes a valid JWT token
	tokenString := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	if tokenString == "" {
		http.Error(w, "Unauthorized API token", http.StatusUnauthorized)
		return
	}
	ctx := context.WithValue(r.Context(), TokenKey, tokenString)

	req := ReqEvent{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
	}
	event := domain.Event(req.Event)
	res, err := e.EventUseCase.Event(ctx, &event)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		if err := json.NewEncoder(w).Encode("Invalid Token"); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (e *EventHandler) EventList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	ctx := r.Context()
	events := &domain.EventCriteria{}
	eventList, err := e.EventUseCase.EventList(ctx, events)
	if err != nil {
		log.Println(err)
	}
	er := json.NewEncoder(w).Encode(eventList)
	if err != nil {
		log.Println(er)
	}
}

func (e *EventHandler) MyEventList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	// Check that the request includes a valid JWT token
	tokenString := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	if tokenString == "" {
		http.Error(w, "Unauthorized API token", http.StatusUnauthorized)
		return
	}

	ctx := context.WithValue(r.Context(), TokenKey, tokenString)

	events := &domain.EventCriteria{}
	eventList, err := e.EventUseCase.MyEventList(ctx, events)
	if err != nil {
		log.Println(err)
	}
	er := json.NewEncoder(w).Encode(eventList)
	if err != nil {
		log.Println(er)
	}
}
