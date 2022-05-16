package httphandler

import "github.com/isurusiri/monorepo-microservices/insight/model"



type (
	// InsightsDTO for get - /insights
	InsightsDTO struct {
		Data []model.Insight `json:"data"`
	}

	// InsightDTO for get - /insight
	InsightDTO struct {
		Data model.Insight `json:"data"`
	}
)
