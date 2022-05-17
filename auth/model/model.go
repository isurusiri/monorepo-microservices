package model

import "gopkg.in/mgo.v2/bson"

type (
	// User model
	User struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName string        `json:"firstname"`
		LastName  string        `json:"lastname"`
		Username  string        `json:"username"`
		Password  string        `json:"password"`
	}

	// AuthToken model
	AuthToken struct {
		ID  bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Jti string        `json:"jti"`
	}

	// RefreshToken model
	RefreshToken struct {
		ID    bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Token string        `json:"token"`
	}

	// AuthResponse model
	AuthResponse struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Username  string        `json:"username"`
		CSRFToken string        `json:"csrf_token"`
	}
)
