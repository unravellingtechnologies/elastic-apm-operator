package handlers

import "net/http"

// Setup function prepares
func Setup() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthz())

	return mux
}
