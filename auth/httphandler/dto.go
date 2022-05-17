package httphandler

import "github.com/isurusiri/monorepo-microservices/auth/model"

type (
	// UserDTO is for Post/Put - /users
	UserDTO struct {
		Data model.User `json:"data"`
	}
)
