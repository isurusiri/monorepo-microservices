package storage

import "github.com/isurusiri/monorepo-microservices/user/model"

// Storage is the interface defining data access contract for Storage model
type Storage interface {
	GetAll() []model.User
	Create(user *model.User) error
	Delete(id string) error
	Ping() error
}
