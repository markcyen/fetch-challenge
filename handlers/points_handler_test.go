package handlers

import (
	"fetch-challenge/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestGetPointsHandler(t *testing.T) {
	// Backup the original `receipts`
	originalReceipts := receipts

	// Restore the original `receipts` after tests
	defer func() {
		receipts = originalReceipts
	}()

	// Define mock receipt data
	mockReceipts := map[string]models.Receipt{
		"123456789": {
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Items: []models.Item{
				{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
				{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
				{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
				{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
			},
			Total:        "35.35",
		},
	}

	tests := []struct {
		name           string
		id             string
		expectedStatus int
		expectedBody   string
		mockReceipts   map[string]models.Receipt
	}{
		{
			name:           "Valid ID",
			id:             "123456789",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"points":28}`, // Update with expected points
			mockReceipts:   mockReceipts,
		},
		{
			name:           "Non-Existent ID",
			id:             "nonexistent-id",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Receipt ID not found\n",
			mockReceipts:   mockReceipts,
		},
		{
			name:           "Missing ID",
			id:             "",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Missing receipt ID\n",
			mockReceipts:   mockReceipts,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Override `receipts` for the test case
			receipts = tc.mockReceipts

			req, err := http.NewRequest(http.MethodGet, "/receipts/"+tc.id+"/points", nil)
			require.NoError(t, err)

			// Mock the path variables for the request
			req = mux.SetURLVars(req, map[string]string{"id": tc.id})

			rr := httptest.NewRecorder()
			GetPointsHandler(rr, req)

			require.Equal(t, tc.expectedStatus, rr.Code)
			require.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rr.Body.String()))
		})
	}
}
