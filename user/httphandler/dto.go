package httphandler

import "github.com/isurusiri/monorepo-microservices/user/model"

type (
	// UsersDTO for Get - /user
	UsersDTO struct {
		Data []model.User `json:"data"`
	}

	// UserDTO for Post/Put - /users
	UserDTO struct {
		Data model.User `json:"data"`
	}
)
