package apis

import (
	"encoding/json"
	"fmt"
	"ivai-api/models"
	"net/http"
)

func GetTemplateFilters() http.HandlerFunc {
	route := "/api/templates/filters"

	return func(w http.ResponseWriter, r *http.Request) {
		filters, err := models.GetTemplateFilters()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonBytes, err := json.Marshal(filters)
		if err != nil {
			fmt.Println("Error marshal JSON:", err)
			return
		}

		var data any
		err = json.Unmarshal(jsonBytes, &data)
		if err != nil {
			fmt.Println("Error Unmarshal Bytes:", err)
			return
		}

		RespondWith(w, r, route, Response{
			Success: true,
			Message: "OK",
			Data:    data,
		})
	}
}
