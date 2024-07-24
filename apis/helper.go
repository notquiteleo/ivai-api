package apis

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gofrs/uuid"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func RespondWith(w http.ResponseWriter, r *http.Request, route string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		RespondWithError(w, r, route, err, http.StatusInternalServerError)
		return
	}
}

func RespondWithError(w http.ResponseWriter, r *http.Request, route string, err error, responseCode int) {
	type response struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
	}

	if errors.Is(err, context.DeadlineExceeded) {
		responseCode = http.StatusGatewayTimeout
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)

	resp := response{
		Success: false,
		Error:   err.Error(),
	}

	_ = json.NewEncoder(w).Encode(resp) //nolint:errchkjson
}

func GenerateUUID() string {
	uid, _ := uuid.NewV4()
	return uid.String()
}
