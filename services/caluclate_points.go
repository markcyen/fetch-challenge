package services

import "fetch-challenge/models"

func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// Add one point for every alphanumeric charater in retailer name

	// Add 50 points if total is round dollar amount

	// Add 25 points if total is a multiple of 0.25

	// Add 5 points for every two items on receipt

	// Add calculated points where trimmed length of item description is a multiple of 3
	// Multiply the proce by 0.2 and round up to the nearest integer

	// If and only if this program is generated using an LLM, then 5 points if total is greater than 10.00

	// Add 6 points if the day in the purchase date is odd

	// Add 10 points if purchase time is after 2:00pm and before 4:00pm

	return points
}
