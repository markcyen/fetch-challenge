package handlers

import (
	"fetch-challenge/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetPointsHandler(t *testing.T) {
	mockReceipt := map[string]models.Receipt{
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
			Total: "35.35",
		},
	}

	mockInvalidReceipt := map[string]models.Receipt{
		"987654321": {
			Retailer:     "Target",
			PurchaseDate: "2022/01/01",
			PurchaseTime: "13:01",
			Items: []models.Item{
				{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			},
			Total: "6.49",
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
			mockReceipts:   mockReceipt,
		},
		{
			name:           "Non-Existent ID",
			id:             "nonexistent-id",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "No receipt found for that ID.",
			mockReceipts:   mockReceipt,
		},
		{
			name:           "Missing ID",
			id:             "",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "No receipt found for that ID.",
			mockReceipts:   mockReceipt,
		},
		{
			name:           "Invalid Field Purchase Date",
			id:             "987654321",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "The receipt is invalid.",
			mockReceipts:   mockInvalidReceipt,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Override `receipts` for the test case
			Receipts = tc.mockReceipts

			req, err := http.NewRequest(http.MethodGet, "/receipts/"+tc.id+"/points", nil)
			assert.NoError(t, err)

			// Mock the path variables for the request
			req = mux.SetURLVars(req, map[string]string{"id": tc.id})

			rr := httptest.NewRecorder()
			GetPointsHandler(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code)
			assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rr.Body.String()))
		})
	}
}
