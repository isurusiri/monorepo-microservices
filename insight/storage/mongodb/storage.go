package mongodb

import (
	"github.com/isurusiri/monorepo-microservices/insight/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Storage holds the database connection
type Storage struct {
	*mgo.Database
}

// GetAll queries and return all Insight documents from the collection.
func (s *Storage) GetAll() []model.Insight {
	c := s.C("insights")

	var insights []model.Insight
	iter := c.Find(nil).Iter()
	result := model.Insight{}
	for iter.Next(&result) {
		insights = append(insights, result)
	}
	return insights
}

// Create inserts a new document to the insights collection.
func (s *Storage) Create(insight *model.Insight) error {
	c := s.C("insight")

	objID := bson.NewObjectId()
	insight.ID = objID
	err := c.Insert(&insight)

	return err
}

// Delete removes a document identified by id.
func (s *Storage) Delete(id string) error {
	c := s.C("insight")

	err := c.Remove((bson.M{"_id": bson.ObjectIdHex(id)}))
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
