package services

import (
	"fetch-challenge/models"
	"testing"

	"github.com/stretchr/testify/require"
)

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
			name: "One Point for Retailer Name",
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
			name: "50 Points for Round Total and Multiple of 0.25",
			receipt: models.Receipt{
				Retailer:     "",
				PurchaseDate: "2023-07-14",
				PurchaseTime: "12:30",
				Total:        "1.00",
				Items:        []models.Item{
					{ShortDescription: "Water", Price: "1.00"},
				},
			},
			expectedPoints: 75,
			expectError:    false,
		},
		{
			name: "5 Points for Every Two Items",
			receipt: models.Receipt{
				Retailer:     "",
				PurchaseDate: "2023-07-14",
				PurchaseTime: "13:30",
				Total:        "4.24",
				Items: []models.Item{
					{ShortDescription: "Coke", Price: "2.23"},
					{ShortDescription: "Water", Price: "2.01"},
				},
			},
			expectedPoints: 5,
			expectError:    false,
		},
		{
			name: "Trimmed Length as Multiple of 3",
			receipt: models.Receipt{
				Retailer:     "",
				PurchaseDate: "2023-07-14",
				PurchaseTime: "12:30",
				Total:        "7.01",
				Items: []models.Item{
					{ShortDescription: "Coca-Cola", Price: "5.00"},
				},
			},
			expectedPoints: 1,
			expectError:    false,
		},
		{
			name: "Purchased on Odd Date",
			receipt: models.Receipt{
				Retailer:     "",
				PurchaseDate: "2023-07-13",
				PurchaseTime: "12:30",
				Total:        "2.01",
				Items:        []models.Item{
					{ShortDescription: "Water", Price: "2.01"},
				},
			},
			expectedPoints: 6,
			expectError:    false,
		},
		{
			name: "Purchased Between 2:00pm and 4:00pm",
			receipt: models.Receipt{
				Retailer:     "",
				PurchaseDate: "2023-07-14",
				PurchaseTime: "14:01",
				Total:        "1.27",
				Items: []models.Item{
					{ShortDescription: "Water", Price: "1.27"},
				},
			},
			expectedPoints: 10,
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
