package http

import (
	"encoding/json"
	"net/http"
)

type serviceUpCheck func() bool

// NewHealthcheckHandler creates a HTTP handler for serving a healthcheck endpoint.
func NewHealthcheckHandler(dbServiceAlive serviceUpCheck) http.HandlerFunc {
	type HealthcheckResponse struct {
		Status string `json:"status"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		status := "fail"
		if dbServiceAlive() {
			status = "ok"
		}
		response := HealthcheckResponse{status}
		responseBytes, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBytes)
	}
}
