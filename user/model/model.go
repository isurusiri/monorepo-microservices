package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// User represents the model of user document.
	User struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name      string        `json:"name"`
		Type      string        `json:"type"`
		Email     string        `json:"email"`
		Username  string        `json:"username"`
		CreatedOn time.Time     `json:"createdOn"`
	}
)
