# fetch-challenge

## Description

Fetch Backend Code Challenge: [receipt-processor-challenge](https://github.com/fetch-rewards/receipt-processor-challenge)


## Language

This program was accomplished in Go 1.23.

## Dependencies

- gorilla/mux for in-memory storage
- google/uuid for id pattern
- regexp for setting pattern on valid fields (ie, retailer name, date, time, total, description and price)
- crypto/sha256 for hashing to identify existing receipts
- stretchr/testify for standard Go testing

## Running the program

For `POST /receipts/process`:

In the terminal, type in the command `go run main.go`. In a separate terminal tab, curl the following -
```bash
curl -X POST http://localhost:8080/receipts/process \             
-H "Content-Type: application/json" \
-d '{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}'

```
Response should be the following where ID is a generated unique ID for the receipt data that was submitted: 
```bash
{"id":"<UNIQUE_ID>"}
```

For `GET /receipts/{id}/points`:
In the terminal, curl the following command using the response from the POST above:
```bash
curl -X GET http://localhost:8080/receipts/<UNIQUE_ID>/points
```
Response should be the points calculated based on the [Rules](https://github.com/fetch-rewards/receipt-processor-challenge/tree/main?tab=readme-ov-file#rules):
```bash
{"points":109}
```

## Testing

To run test files, run `go test ./...` in the command line to see results. 

## About the author


