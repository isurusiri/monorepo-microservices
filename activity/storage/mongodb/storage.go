package mongodb

import (
	"github.com/isurusiri/monorepo-microservices/activity/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Storage holds the database connection.
type Storage struct {
	*mgo.Database
}

// GetAll queries and return all Activity documents from the database.
func (s *Storage) GetAll() []model.Activity {
	c := s.C("activities")

	var activities []model.Activity
	iter := c.Find(nil).Iter()
	result := model.Activity{}
	for iter.Next(&result) {
		activities = append(activities, result)
	}
	return activities
}

// Create inserts a new entity in to the activities collection.
func (s *Storage) Create(activity *model.Activity) error {
	c := s.C("activities")

	objctID := bson.NewObjectId()
	activity.ID = objctID
	err := c.Insert(&activity)

	return err
}

// Delete removes a document identified by id from activities
// collection.
func (s *Storage) Delete(id string) error {
	c := s.C("activities")

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
