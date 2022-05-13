package httphandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/isurusiri/monorepo-microservices/activity/storage"
	"gopkg.in/mgo.v2"
)

// GetActivities is the handler for HTTP Get - /activities
// Get all activities
func GetActivities(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		activities := s.GetAll()

		// Create response object
		j, err := json.Marshal(ActivitiesDTO{Data: activities})
		if err != nil {
			panic(err)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(j)
	})
}

// CreateActivity is the handler for HTTP Post - /activities
// Create a new activity document
func CreateActivity(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var activityDTO ActivityDTO
		err := json.NewDecoder(r.Body).Decode(&activityDTO)
		if err != nil {
			panic(err)
		}
		activity := activityDTO.Data

		s.Create(&activity)

		j, err := json.Marshal(activityDTO)
		if err != nil {
			panic(err)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		rw.Write(j)
	})
}

// DeleteActivity is the handler for HTTP Delete - /activities
// Delete an activity document
func DeleteActivity(s storage.Storage) http.HandlerFunc {
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
