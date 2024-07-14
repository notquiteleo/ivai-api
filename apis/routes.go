package apis

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// r.Use(middleware.Heartbeat(s.getPath("/health")))
	// TODO add middleware to check if user is logged in

	r.Get("/api/users/verity_code", VerityCode())
	r.Get("/api/users/mobile_login", MobileLogin())
	r.Get("/api/users/wx_login", WxLogin())

	r.Get("/api/templates/filters", GetTemplateFilters())
	r.Get("/api/templates", GetTemplates())

	r.Delete("/api/resumes/{id}", DeleteResume())

	return r
}
