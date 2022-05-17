package httphandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/isurusiri/monorepo-microservices/user/storage"
	"gopkg.in/mgo.v2"
)

// GetUsers is the handler for HTTP Get - /users
// Get all users
func GetUsers(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		users := s.GetAll()

		j, err := json.Marshal(UsersDTO{Data: users})
		if err != nil {
			panic(nil)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(j)
	})
}

// CreateUser is the handler for HTTP Post - /users
// Create a new user document
func CreateUser(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var userDTO UserDTO
		err := json.NewDecoder(r.Body).Decode(&userDTO)
		if err != nil {
			panic(err)
		}
		user := userDTO.Data

		s.Create(&user)

		j, err := json.Marshal(userDTO)
		if err != nil {
			panic(err)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		rw.Write(j)
	})
}

// DeleteUser is the handler for HTTP Delete - /users
// Delete an Users document
func DeleteUser(s storage.Storage) http.HandlerFunc {
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

// GetReadiness indicates if connection is ready.
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

// GetLiveness indicates service availability.
func GetLiveness() http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("ok"))
	})
}
