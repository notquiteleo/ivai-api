package apis

import (
	"ivai-api/middlewares/jwt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// 配置 CORS 中间件
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"}, // 允许的域名
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // 允许预检请求的缓存时间
	}
	r.Use(cors.Handler(corsOptions))

	r.Group(func(r chi.Router) {
		r.Get("/api/users/verity_code", VerityCode())
		r.Post("/api/users/mobile_login", MobileLogin())
		r.Get("/api/users/wx_login", WxLogin())
	})

	r.Group(func(r chi.Router) {
		r.Use(jwt.JWTMiddleware())

		r.Group(func(r chi.Router) {
			r.Get("/api/templates/filters", GetTemplateFilters())
			r.Get("/api/templates", GetTemplates())
			r.Post("/api/templates", CreateTemplates())
		})

		r.Group(func(r chi.Router) {
			r.Get("/api/resumes", GetResumes())
			r.Get("/api/resumes/{id}", GetResume())
			r.Post("/api/resumes/upload", UploadResume())
			r.Post("/api/resumes", CreateResume())
			r.Put("/api/resumes/{id}", UpdateResume())
			r.Delete("/api/resumes/{id}", DeleteResume())
		})
	})

	return r
}
