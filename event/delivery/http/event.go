package http

import (
	"encoding/json"
	"github.com/event_bright/domain"
	"github.com/go-chi/chi"
	"log"
	"net/http"
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
	req := ReqEvent{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
	}

	ctx := r.Context()
	event := domain.Event(req.Event)
	res, _ := e.EventUseCase.Event(ctx, &event)

	er := json.NewEncoder(w).Encode(res)
	if er != nil {
		log.Println(er)
	}
}
