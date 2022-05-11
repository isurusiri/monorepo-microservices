package httphandler

import "github.com/isurusiri/monorepo-microservices/behaviour/model"

type (
	// BehavioursDTO for Get - /behaviours
	BehavioursDTO struct {
		Data []model.Behaviour `json:"data"`
	}

	// BehaviourDTO for Post/Put - /behaviours
	BehaviourDTO struct {
		Data model.Behaviour `json:"data"`
	}
)
