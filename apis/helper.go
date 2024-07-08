package apis

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type Response struct {
	Success bool
	Message string
	Data    any
}

func RespondWith(w http.ResponseWriter, r *http.Request, route string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		RespondWithError(w, r, route, err, http.StatusInternalServerError)
		return
	}

	// IncrementRequestsTotal(r.Method, route, http.StatusOK)
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

	// IncrementRequestsTotal(r.Method, route, responseCode)
}

// TODO prometheus
// func IncrementRequestsTotal(method, route string, responseCode int) {
// 	MetricsRequestsTotal.With(prometheus.Labels{"code": strconv.Itoa(responseCode), "method": method, "route": route}).Inc()
// }
