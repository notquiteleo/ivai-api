package apis

import (
	"ivai-api/models"
	"net/http"
	"strconv"
)

func GetTemplates() http.HandlerFunc {
	route := "/api/templates"

	return func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("template_type")
		templateType, err := strconv.ParseInt(param, 10, 32)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		templates, err := models.GetTemplates(templateType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		TestRenderSimpleTemplate(w,r)

		RespondWith(w, r, route, Response{
			Success: true,
			Message: "OK",
			Data:    templates,
		})
	}
}

// create single template
// 1. first, receive origin template from front end
// 2. second, parse template to html content by liquid and split div block to create template and modules
// 3. third, render html content to get preview image by mock resume json data

func CreateTemplates() http.HandlerFunc {
	route := "/api/templates"

	return func(w http.ResponseWriter, r *http.Request) {
		RespondWith(w, r, route, Response{
			Success: true,
			Message: "OK",
		})
	}
}