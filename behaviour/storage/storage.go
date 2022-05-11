package storage

import "github.com/isurusiri/monorepo-microservices/behaviour/model"

// Storage is the interface defining data access contract for
// Storage model
type Storage interface {
	GetAll() []model.Behaviour
	Create(behaviour *model.Behaviour) error
	Delete(id string) error
	Ping() error
}
