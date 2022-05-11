package mongodb

import (
	"github.com/isurusiri/monorepo-microservices/behaviour/model"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Storage holds the database connection.
type Storage struct {
	*mgo.Database
}

// GetAll queries and return all Behaviour entities from the database.
func (s *Storage) GetAll() []model.Behaviour {
	c := s.C("behaviours")

	var behaviours []model.Behaviour
	iter := c.Find(nil).Iter()
	result := model.Behaviour{}
	for iter.Next(&result) {
		behaviours = append(behaviours, result)
	}
	return behaviours
}

// Create inserts a new entity in to behaviours collection.
func (s *Storage) Create(behaviour *model.Behaviour) error {
	c := s.C("behaviours")

	objID := bson.NewObjectId()
	behaviour.ID = objID
	err := c.Insert(&behaviour)

	return err
}

// Delete removes an entity record identified by id.
func (s *Storage) Delete(id string) error {
	c := s.C("behaviours")

	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

// Ping checks the db connectivity
func (s *Storage) Ping() error {
	var pingStatus error

	err := s.Session.Ping()
	if err != nil {
		pingStatus = err
	} else {
		pingStatus = nil
	}

	return pingStatus
}
