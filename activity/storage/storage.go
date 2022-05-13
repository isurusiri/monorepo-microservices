package storage

import "github.com/isurusiri/monorepo-microservices/activity/model"

// Storage is the interface defining data access contract for Storage model
type Storage interface {
	GetAll() []model.Activity
	Create(activity *model.Activity) error
	Delete(id string) error
	Ping() error
}
