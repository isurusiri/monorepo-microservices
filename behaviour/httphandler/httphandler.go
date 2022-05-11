package httphandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/isurusiri/monorepo-microservices/behaviour/storage"
	"gopkg.in/mgo.v2"
)

// GetBehaviours is the Handler for HTTP Get - /behaviours
// Get all Behaviours
func GetBehaviours(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		behaviours := s.GetAll()
	
		// Create response object
		j, err := json.Marshal(BehavioursDTO{Data: behaviours})
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})
}

// CreateBehaviour is the handler for HTTP Post - /behaviours
// Create a new behaviour document
func CreateBehaviour(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var behaviourDTO BehaviourDTO
		// decode incoming dto object
		err := json.NewDecoder(r.Body).Decode(&behaviourDTO)
		if err != nil {
			panic(err)
		}
		behaviour := &behaviourDTO.Data

		s.Create(behaviour)

		j, err := json.Marshal(behaviourDTO)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	})
}

// DeleteBehaviour is the handler for HTTP Delete - /behaviours
// Delete a behaviour document
func DeleteBehaviour(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		err := s.Delete(id)
		if err != nil {
			if err == mgo.ErrNotFound {
				rw.WriteHeader(http.StatusNotFound)
				return
			} else {
				panic(err)
			}
		}

		rw.WriteHeader(http.StatusNoContent)
	})
}

// GetReadiness indicates if connection is ready
func GetReadiness(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		err := s.Ping()
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(fmt.Sprintf("error: %v", err)))
		} else {
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte("ok"))
		}
	})
}

// GetLiveness indicates service availability
func GetLiveness() http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("ok"))
	})
}
