package apis

import (
	"ivai-api/middlewares/jwt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/api/users/verity_code", VerityCode())
	r.Post("/api/users/mobile_login", MobileLogin())
	r.Get("/api/users/wx_login", WxLogin())

	r.Group(func(r chi.Router) {
		r.Use(jwt.JWTMiddleware())
		
		r.Get("/api/templates/filters", GetTemplateFilters())
		r.Get("/api/templates", GetTemplates())

		r.Delete("/api/resumes/{id}", DeleteResume())
	})

	return r
}
