package models

import (
	"time"

	"github.com/google/uuid"
)

type Repository struct {
	BaseAggregateSourced

	OrganisationId uuid.UUID `json:"organisation_id" format:"uuid" required:"true"`
	Name           string    `json:"name" required:"true"`
	FullName       string    `json:"full_name" required:"true"`
	DefaultBranch  string    `json:"default_branch" required:"true"`
	CreatedAt      time.Time `json:"created_at" required:"true"`
	Url            string    `json:"url" required:"true"`

	ClusterId uuid.UUID `json:"cluster_id" format:"uuid" required:"true"`
	PlanId    uuid.UUID `json:"plan_id" format:"uuid" required:"true"`
}
