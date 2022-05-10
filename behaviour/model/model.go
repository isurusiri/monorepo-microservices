package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// Behaviour represents the model of behaviour database entity.
	Behaviour struct {
		ID         bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name       string        `json:"name"`
		Type       string        `json:"type"`
		Confidence string        `json:"confidence"`
		CreatedOn  time.Time     `json:"createdon,omitempty"`
	}
)
