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

		RespondWith(w, r, route, Response{
			Success: true,
			Message: "OK",
			Data:    templates,
		})
	}
}
