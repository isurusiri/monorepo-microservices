package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// Activity represents the model of activity database entity.
	Activity struct {
		ID         bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name       string        `json:"name"`
		Type       string        `json:"type"`
		DetectedOn time.Time     `json:"detectedOn"`
		CreatedOn  time.Time     `json:"createdOn"`
	}
)
