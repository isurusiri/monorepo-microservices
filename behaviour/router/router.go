package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/isurusiri/monorepo-microservices/behaviour/httphandler"
	"github.com/isurusiri/monorepo-microservices/behaviour/storage"
)

// InitRouter registers routes of the service under /behaviours path.
func InitRouter(r *chi.Mux, s storage.Storage) *chi.Mux {
	r.Route("/behaviours", func(r chi.Router) {
		r.Get("/", httphandler.GetBehaviours(s))
		r.Post("/", httphandler.CreateBehaviour(s))
		r.Delete("/", httphandler.DeleteBehaviour(s))
	})

	r.Get("/health", httphandler.GetReadiness(s))
	r.Get("/healthz", httphandler.GetLiveness())

	return r
}
