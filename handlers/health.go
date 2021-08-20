package handlers

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

// healthz function provides a healthcheck endpoint for the webhook
func healthz() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("{\"status\": \"ok\"}"))
		if err != nil {
			log.Error("Not able to answer request", err)
			return
		}
	}
}
