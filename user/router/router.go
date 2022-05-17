package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/isurusiri/monorepo-microservices/user/httphandler"
	"github.com/isurusiri/monorepo-microservices/user/storage"
)

// InitRouter registers routes of the service under /users path.
func InitRouter(r *chi.Mux, s storage.Storage) *chi.Mux {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", httphandler.GetUsers(s))
		r.Post("/", httphandler.CreateUser(s))
		r.Delete("/", httphandler.DeleteUser(s))
	})

	r.Get("/health", httphandler.GetReadiness(s))
	r.Get("/healthz", httphandler.GetLiveness())

	return r
}
