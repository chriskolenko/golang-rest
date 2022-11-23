package handler

import (
	"context"
	"net/http"

	"github.com/contextcloud/graceful/config"

	"github.com/go-chi/chi/middleware"
	"github.com/google/uuid"

	"github.com/swaggest/jsonschema-go"
	"github.com/swaggest/openapi-go/openapi3"
	"github.com/swaggest/rest/web"
	"github.com/swaggest/swgui/v4emb"

	"github.com/chriskolenko/golang-rest/controllers"
)

func NewHandler(ctx context.Context, cfg *config.Config) (http.Handler, error) {
	serviceInfo := openapi3.Info{}
	serviceInfo.
		WithTitle(cfg.ServiceName).
		WithVersion(cfg.Version)

	s := web.DefaultService()
	s.OpenAPI.Info = serviceInfo

	// Create custom schema mapping for 3rd party type.
	uuidDef := jsonschema.Schema{}
	uuidDef.AddType(jsonschema.String)
	uuidDef.WithFormat("uuid")
	uuidDef.WithExamples("248df4b7-aa70-47b8-a036-33ac447e668d")
	s.OpenAPICollector.Reflector().AddTypeMapping(uuid.UUID{}, uuidDef)
	s.OpenAPICollector.Reflector().InlineDefinition(uuid.UUID{})

	s.Use(middleware.Logger)
	s.Docs("/docs", v4emb.New)
	s.Mount("/debug", middleware.Profiler())
	s.Get("/organisations", controllers.GetOrganisations())
	s.Get("/repositories", controllers.GetRepositories())

	return s, nil
}
