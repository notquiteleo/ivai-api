package apis

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// r.Use(middleware.Heartbeat(s.getPath("/health")))

	r.Get("/api/templates/filters", GetTemplateFilters())
	r.Get("/api/templates", GetTemplates())
	// r.Delete("/api/resumes/{id}", DeleteResume())

	return r
}
