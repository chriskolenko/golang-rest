package controllers

import (
	"context"

	"github.com/chriskolenko/golang-rest/models"
	"github.com/google/uuid"
	"github.com/swaggest/usecase"
)

func GetOrganisations() usecase.Interactor {
	type OrganisationQuery struct {
		Namespace string       `header:"X-Namespace" required:"true"`
		Ids       *[]uuid.UUID `query:"ids"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, input OrganisationQuery, output *[]*models.Organisation) error {
		return nil
	})

	u.SetName("GetOrganisations")
	return u
}
