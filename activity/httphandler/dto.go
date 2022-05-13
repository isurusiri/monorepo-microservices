package httphandler

import "github.com/isurusiri/monorepo-microservices/activity/model"

type (
	// ActivitiesDTO for Get - /activities
	ActivitiesDTO struct {
		Data []model.Activity `json:"data"`
	}

	// ActivityDTO for Post/Put - /activities
	ActivityDTO struct {
		Data model.Activity `json:"data"`
	}
)
