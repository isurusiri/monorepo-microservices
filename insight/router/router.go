package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/isurusiri/monorepo-microservices/insight/httphandler"
	"github.com/isurusiri/monorepo-microservices/insight/storage"
)

// InitRouter registers routes of the service under /insights path.
func InitRouter(r *chi.Mux, s storage.Storage) *chi.Mux {
	r.Route("/insights", func(r chi.Router) {
		r.Get("/", httphandler.GetInsights(s))
		r.Post("/", httphandler.CreateInsight(s))
		r.Delete("/", httphandler.DeleteInsight(s))
	})

	r.Get("/health", httphandler.GetReadiness(s))
	r.Get("/healthz", httphandler.GetLiveness())

	return r
}
