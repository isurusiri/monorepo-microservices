package httphandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/isurusiri/monorepo-microservices/auth/httphandler/contract"
	"github.com/isurusiri/monorepo-microservices/auth/storage/mongodb"
)

// AuthHandler holds required connection objects.
type AuthHandler struct {
	authContract contract.AuthenticationContract
	storage      *mongodb.Storage
}

// NewAuthHandler initiate and returns a new auth handler
func NewAuthHandler(authContract contract.AuthenticationContract,
		storage *mongodb.Storage) *AuthHandler {
			return &AuthHandler{
				authContract: authContract,
				storage:      storage,
			}
		}

// Login is the handler for http POST - /login
// return auth token with refresh token.
func (ah *AuthHandler) Login() http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var userDto UserDTO

		err := json.NewDecoder(r.Body).Decode(&userDto)
		if err != nil {
			panic(err)
		}

		token, err := ah.authContract.Login(userDto.Data)
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
		}

		setAuthCookie(&rw, token.AuthToken, token.RefreshToken)
		setRefreshCookie(&rw, token.AuthToken, token.RefreshToken)
		
		rw.Header().Set("X-CSRF-Token", token.CSRFKey)
		rw.WriteHeader(http.StatusOK)
	})
}

// Authenticate is the handler for http POST - /authenticate
// return HTTP ok if token valid and HTTP 400 if not valid.
func (ah *AuthHandler) Authenticate() http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		authCookie, err := r.Cookie("AuthToken")
		if err == http.ErrNoCookie {
			ah.nullifyTokenCookies(&rw, r)
			http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		} else if err != nil {
			ah.nullifyTokenCookies(&rw, r)
			http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		refreshToken, err := r.Cookie("RefreshToken")
		if err == http.ErrNoCookie {
			ah.nullifyTokenCookies(&rw, r)
			http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		} else if err != nil {
			ah.nullifyTokenCookies(&rw, r)
			http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		csrfKey := getCSRFKey(r)

		token, err := ah.authContract.Authenticate(authCookie.Value, refreshToken.Value, csrfKey)
		if err != nil {
			println(err.Error())
			if err.Error() == "Unauthorized" {
				http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			} else {
				http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}

		setAuthCookie(&rw, token.AuthToken, token.RefreshToken)
		setRefreshCookie(&rw, token.AuthToken, token.RefreshToken)
		rw.Header().Set("X-CSRF-Token", token.CSRFKey)
		rw.WriteHeader(http.StatusOK)
	})
}

// GetReadiness indicates if connection is ready.
func (ah *AuthHandler) GetReadiness() http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		err := ah.storage.Ping()
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
func (ah *AuthHandler) GetLiveness() http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("ok"))
	})
}

func getCSRFKey(r *http.Request) string {
	csrfFromFrom := r.FormValue("X-CSRF-Token")

	if csrfFromFrom != "" {
		return csrfFromFrom
	} else {
		return r.Header.Get("X-CSRF-Token")
	}
}

func setAuthCookie(w *http.ResponseWriter, authTokenString string, refreshTokenString string) {
	authCookie := http.Cookie{
		Name:     "AuthToken",
		Value:    authTokenString,
		HttpOnly: true,
	}

	http.SetCookie(*w, &authCookie)
}

func setRefreshCookie(w *http.ResponseWriter, authTokenString string, refreshTokenString string) {
	refreshCookie := http.Cookie{
		Name:     "RefreshToken",
		Value:    refreshTokenString,
		HttpOnly: true,
	}

	http.SetCookie(*w, &refreshCookie)
}

func (ah *AuthHandler) nullifyTokenCookies(w *http.ResponseWriter, r *http.Request) {
	authCookie := http.Cookie{
		Name:    "AuthToken",
		Value:    "",
		Expires:  time.Now().Add(-1000 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(*w, &authCookie)

	refreshCookie := http.Cookie{
		Name: "RefreshToken",
		Value: "",
		Expires: time.Now().Add(-1000 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(*w, &refreshCookie)

	RefreshCookie, refreshErr := r.Cookie("RefreshToken")
	if refreshErr == http.ErrNoCookie {
		return
	} else if refreshErr != nil {
		log.Panic("panic %v", refreshErr)
		http.Error(*w, http.StatusText(500), 500)
	}

	ah.authContract.RevokeRefreshToken(RefreshCookie.Value)
}