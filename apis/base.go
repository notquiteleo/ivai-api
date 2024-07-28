package apis

import (
	"net/http"

	"ivai-api/models"
)

func GetUserFromContext(r *http.Request) *models.Users {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		return nil
	}
	user, err := models.GetUserByConditions(map[string]interface{}{"id": userID})
	if err != nil {
		return nil
	}
	return user
}
