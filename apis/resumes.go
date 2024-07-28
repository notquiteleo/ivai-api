package apis

import (
	"ivai-api/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetResume() http.HandlerFunc {
	route := "/api/resumes/{id}"

	return func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "id")
		resumeID, err := strconv.ParseInt(param, 10, 64)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var resume *models.Resumes
		err = models.DB.Where("id = ?", resumeID).Find(&resume).Error
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

func GetResumes() http.HandlerFunc {
	route := "/api/resumes"

	return func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "id")
		resumeID, err := strconv.ParseInt(param, 10, 64)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user := GetUserFromContext(r)
		if user != nil {
			RespondWith(w, r, route, Response{
				Success: false,
				Message: "user not found",
			})
		}

		var resume *models.Resumes
		err = models.DB.Where("id = ?", resumeID).Find(&resume).Error
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

// TODO AI
func UploadResume() http.HandlerFunc {
	route := "/api/resumes/upload"

	return func(w http.ResponseWriter, r *http.Request) {

		RespondWith(w, r, route, Response{
			Success: true,
			Message: "OK",
		})
	}
}

func CreateResume() http.HandlerFunc {
	route := "/api/resumes"

	return func(w http.ResponseWriter, r *http.Request) {

		RespondWith(w, r, route, Response{
			Success: true,
			Message: "OK",
		})
	}
}

func UpdateResume() http.HandlerFunc {
	route := "/api/resumes"

	return func(w http.ResponseWriter, r *http.Request) {

		RespondWith(w, r, route, Response{
			Success: true,
			Message: "OK",
		})
	}
}

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
