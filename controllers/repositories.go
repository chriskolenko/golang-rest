package controllers

import (
	"context"

	"github.com/chriskolenko/golang-rest/models"
	"github.com/google/uuid"
	"github.com/swaggest/usecase"
)

func GetRepositories() usecase.Interactor {
	type RepositoryQuery struct {
		Namespace      string       `header:"X-Namespace" required:"true"`
		Ids            *[]uuid.UUID `query:"ids"`
		OrganisationId *uuid.UUID   `query:"organisation_id"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, input RepositoryQuery, output *[]*models.Repository) error {
		return nil
	})

	u.SetName("GetRepositories")
	return u
}
