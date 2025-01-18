package services

import (
	"fetch-challenge/models"
	"testing"

	"github.com/stretchr/testify/require"
)

// TODO: Update tests
func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name           string
		receipt        models.Receipt
		expectedPoints int
		expectError    bool
	}{
		{
			name: "Basic Receipt - All Scenarios",
			receipt: models.Receipt{
				Retailer:     "B",
				PurchaseDate: "2023-07-13",
				PurchaseTime: "15:30",
				Total:        "10.00",
				Items: []models.Item{
					{ShortDescription: "Coca-Cola", Price: "5.00"},
					{ShortDescription: "Water", Price: "5.00"},
				},
			},
			expectedPoints: 1 + 50 + 25 + 5 + 1 + 6 + 10,
			expectError:    false,
		},
		{
			name: "Invalid Total Format",
			receipt: models.Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2023-07-14",
				PurchaseTime: "15:30",
				Total:        "invalid",
				Items:        nil,
			},
			expectedPoints: 0,
			expectError:    true,
		},
		{
			name: "One point for Retailer Name with one character",
			receipt: models.Receipt{
				Retailer:     "B",
				PurchaseDate: "2023-07-14",
				PurchaseTime: "12:30",
				Total:        "1.29",
				Items:        []models.Item{
					{ShortDescription: "Water", Price: "1.29"},
				},
			},
			expectedPoints: 1,
			expectError:    false,
		},
		{
			name: "Odd Purchase Day",
			receipt: models.Receipt{
				Retailer:     "OddDayRetailer",
				PurchaseDate: "2023-07-15",
				PurchaseTime: "13:00",
				Total:        "10.00",
				Items:        nil,
			},
			expectedPoints: 6,
			expectError:    false,
		},
		{
			name: "Purchase Time in Range",
			receipt: models.Receipt{
				Retailer:     "TimedRetailer",
				PurchaseDate: "2023-07-14",
				PurchaseTime: "14:30",
				Total:        "10.00",
				Items:        nil,
			},
			expectedPoints: 10,
			expectError:    false,
		},
		{
			name: "Item Description Length Points",
			receipt: models.Receipt{
				Retailer:     "Retailer",
				PurchaseDate: "2023-07-14",
				PurchaseTime: "13:00",
				Total:        "10.00",
				Items: []models.Item{
					{ShortDescription: "abc", Price: "1.00"},
				},
			},
			expectedPoints: 1, // "abc" length is 3, which is a multiple of 3
			expectError:    false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			points, err := CalculatePoints(tc.receipt)
			if tc.expectError {
				require.Error(t, err, "Expected an error but didn't get one")
			} else {
				require.NoError(t, err, "Expected no error but got one")
				require.Equal(t, tc.expectedPoints, points, "Points mismatch")
			}
		})
	}
}
