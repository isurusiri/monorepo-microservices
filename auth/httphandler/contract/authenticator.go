package contract

import "github.com/isurusiri/monorepo-microservices/auth/model"

// AuthenticationContract represents the iterface for authentication
type AuthenticationContract interface {
	Login(user model.User) (model.Token, error)
	Authenticate(authToken, refreshToken, csrfToken string) (model.Token, error)
	RevokeRefreshToken(refreshToken string) error
}
