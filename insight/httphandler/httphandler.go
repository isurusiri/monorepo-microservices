package httphandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/isurusiri/monorepo-microservices/insight/storage"
	"gopkg.in/mgo.v2"
)

// GetInsights is the Handler for HTTP Get - /insights
// Get all insights
func GetInsights(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		insights := s.GetAll()

		j, err := json.Marshal(InsightsDTO{Data: insights})
		if err != nil {
			panic(err)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(j)
	})
}

// CreateInsight is the handler for HTTP Post - /insights
// Create a new insight document
func CreateInsight(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var insightDTO InsightDTO

		err := json.NewDecoder(r.Body).Decode(&insightDTO)
		if err != nil {
			panic(err)
		}
		insight := insightDTO.Data

		s.Create(&insight)

		j, err := json.Marshal(insightDTO)
		if err != nil {
			panic(err)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(j)
	})
}

// DeleteInsight is the handler for HTTP Delete - /insights
// Delete an insight document
func DeleteInsight(s storage.Storage) http.HandlerFunc {
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
