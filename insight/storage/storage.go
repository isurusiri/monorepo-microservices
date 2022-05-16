package storage

import "github.com/isurusiri/monorepo-microservices/insight/model"

// Storage is the interface defining data access contract
// for Storage model.
type Storage interface {
	GetAll() []model.Insight
	Create(insight *model.Insight) error
	Delete(id string) error
	Ping() error
}
