package mongodb

import (
	"github.com/isurusiri/monorepo-microservices/user/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Storage holds the database connection.
type Storage struct {
	*mgo.Database
}

// GetAll queries and return all User documents from the database.
func (s *Storage) GetAll() []model.User {
	c := s.C("users")

	var users []model.User
	iter := c.Find(nil).Iter()
	result := model.User{}
	for iter.Next(&result) {
		users = append(users, result)
	}
	return users
}

// Create inserts a new entity in to the users colleaction.
func (s *Storage) Create(user *model.User) error {
	c := s.C("users")

	objectID := bson.NewObjectId()
	user.ID = objectID
	err := c.Insert(user)

	return err
}

// Delete removes a document identified by id from users collection.
func (s *Storage) Delete(id string) error {
	c := s.C("users")

	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

// Ping checks the db connectivity.
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
