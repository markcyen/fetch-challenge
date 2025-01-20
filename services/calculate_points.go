package services

import (
	"fetch-challenge/models"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// CalculatePoints calculate points based on specific criteria in the receipt
func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// Add one point for every alphanumeric character in retailer name
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points++
		}
	}

	total, _ := strconv.ParseFloat(receipt.Total, 64)

	// Add 50 points if total is round dollar amount with no cents
	if total == float64(int(total)) {
		points += 50
	}

	// Add 25 points if total is a multiple of 0.25
	if total == float64(int(total)) {
		points += 25
	}

	// Add 5 points for every two items on receipt
	points += len(receipt.Items) / 2 * 5

	// Add calculated points where trimmed length of item description is a multiple of 3
	// Multiply the price by 0.2 and round up to the nearest integer
	for _, item := range receipt.Items {
		if item.ShortDescription == "" {
			continue
		}
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		cleanDescription := strings.Join(strings.Fields(trimmedDescription), " ")
		if len(cleanDescription)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)

			points += int(math.Ceil(price * 0.2))
		}
	}

	// Add 6 points if the day in the purchase date is odd
	day := receipt.PurchaseDate[len(receipt.PurchaseDate)-2:]
	convertDay, _ := strconv.Atoi(day)
	if convertDay%2 != 0 {
		points += 6
	}

	// Add 10 points if purchase time is after 2:00pm (14:00) and before 4:00pm (16:00)
	startTime, _ := time.Parse("15:04", "14:00")
	endTime, _ := time.Parse("15:04", "16:00")
	convertTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if convertTime.After(startTime) && convertTime.Before(endTime) {
		points += 10
	}

	return points
}
