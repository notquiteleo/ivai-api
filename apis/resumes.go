package apis

import (
	"ivai-api/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func DeleteResume() http.HandlerFunc {
	route := "/api/resumes/{id}"

	return func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "id")
		resumeID, err := strconv.ParseInt(param, 10, 64)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = models.DB.Delete(&models.Resumes{}, resumeID).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		RespondWith(w, r, route, Response{
			Success: true,
			Message: "OK",
		})
	}
}
