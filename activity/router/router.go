package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/isurusiri/monorepo-microservices/activity/httphandler"
	"github.com/isurusiri/monorepo-microservices/activity/storage"
)

// InitRouter registers routes of the service under /activities path.
func InitRouter(r *chi.Mux, s storage.Storage) *chi.Mux {
	r.Route("/activities", func(r chi.Router) {
		r.Get("/", httphandler.GetActivities(s))
		r.Post("/", httphandler.CreateActivity(s))
		r.Delete("/", httphandler.DeleteActivity(s))
	})

	r.Get("/health", httphandler.GetReadiness(s))
	r.Get("/healthz", httphandler.GetLiveness())

	return r
}
