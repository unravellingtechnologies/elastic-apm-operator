package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHealthz function tests the function healthz() to make the healthcheck replies with what we expect
func TestHealthz(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := healthz()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v wanted %v",
			status, http.StatusOK)
	}

	expected := `{"status": "ok"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v wanted %v",
			rr.Body.String(), expected)
	}
}
