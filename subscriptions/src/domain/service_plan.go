package domain

import "github.com/google/uuid"

type ServicePlan struct {
	Id       uuid.UUID
	Name     string
	Features []FeatureOption
}
