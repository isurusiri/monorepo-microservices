package mongodb

import (
	"github.com/isurusiri/monorepo-microservices/auth/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Storage represents the database connection.
type Storage struct {
	*mgo.Database
}

// GetByeUsernameAndPassword queries a user by username.
func (s *Storage) GetByeUsernameAndPassword(username, password string) (model.User, error) {
	c := s.C("users")
	filter := bson.M{"username": username}

	user := model.User{}
	err := c.Find(filter).One(&user)
	return user, err
}

// GetRefreshToken query and return a refresh token.
func (s *Storage) GetRefreshToken(token string) (model.RefreshToken, error) {
	c := s.C("refreshtoken")
	filter := bson.M{"token": token}

	refreshToken := model.RefreshToken{}
	err := c.Find(filter).One(&refreshToken)
	return refreshToken, err
}

// StoreRefreshToken store a new refresh token.
func (s *Storage) StoreRefreshToken(token model.RefreshToken) error {
	c := s.C("refreshtoken")

	existingToken, err := s.GetRefreshToken(token.Token)
	if err != mgo.ErrNotFound {
		return err
	}

	if existingToken.ID != "" {
		return nil
	}

	objID := bson.NewObjectId()
	token.ID = objID
	return c.Insert(&token)
}

// DeleteRefreshToken removes a token from the collection.
func (s *Storage) DeleteRefreshToken(token string) error {
	c := s.C("authtoken")

	filter := bson.M{"token": token}
	err := c.Remove(filter)
	return err
}

// Ping checks if the db connection is alive.
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