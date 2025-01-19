package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	router := mux.NewRouter()

	router.HandleFunc("/receipts/process", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("POST")

	router.HandleFunc("/receipts/{id}/points", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
	}{
		{
			name:           "Valid POST /receipts/process",
			method:         "POST",
			path:           "/receipts/process",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid Method for /receipts/process",
			method:         "GET",
			path:           "/receipts/process",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Valid GET /receipts/{id}/points",
			method:         "GET",
			path:           "/receipts/123/points",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid Path",
			method:         "GET",
			path:           "/invalid/path",
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, tc.path, nil)
			assert.NoError(t, err, "Failed to create request")

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code, "Unexpected status code for %s", tc.name)
		})
	}
}
