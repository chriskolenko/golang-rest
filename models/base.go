package models

import (
	"github.com/google/uuid"
)

type BaseAggregateSourced struct {
	Id        uuid.UUID `json:"id" format:"uuid" required:"true"`
	Namespace string    `json:"namespace" required:"true"`
	Version   int       `json:"version"`
}
