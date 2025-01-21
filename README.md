# fetch-challenge

## Description

Fetch Backend Code Challenge: [receipt-processor-challenge](https://github.com/fetch-rewards/receipt-processor-challenge)


## Language

This program was accomplished in Go 1.23.

## Dependencies

- `gorilla/mux` for in-memory storage
- `google/uuid` for id pattern (pattern from [from api.yml](https://github.com/fetch-rewards/receipt-processor-challenge/blob/main/api.yml))
- `regexp` for setting pattern on field values (patterns from [from api.yml](https://github.com/fetch-rewards/receipt-processor-challenge/blob/main/api.yml))
- `stretchr/testify` for standard Go testing

## Running the program

#### For `POST /receipts/process`:

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
Example `"id"` should look like this:
```bash
{"id":"4d1ebc47-62a8-4116-9c61-ad60d7a343fa"}
```
Note: Feel free to manipulate the field values to check what points are calculated or what values would make the receipt invalid.

#### For `GET /receipts/{id}/points`:
In the terminal, curl the following command using the response from the POST above:
```bash
curl -X GET http://localhost:8080/receipts/<UNIQUE_ID>/points
```
Example GET curl:
```bash
curl -X GET http://localhost:8080/receipts/4d1ebc47-62a8-4116-9c61-ad60d7a343fa/points
```
Response should be the points calculated based on the [Rules](https://github.com/fetch-rewards/receipt-processor-challenge/tree/main?tab=readme-ov-file#rules):
```bash
{"points":109}
```

## Testing

To run test files and get the test coverage, type in the following commands:
```bash
go test -cover ./...
```
Results from the test coverage:
```bash
ok  	fetch-challenge	(cached)	coverage: 50.0% of statements
?   	fetch-challenge/models	[no test files]
ok  	fetch-challenge/handlers	(cached)	coverage: 95.9% of statements
ok  	fetch-challenge/services	(cached)	coverage: 96.9% of statements
```

## About the author

Hi 👋🏼 I am Mark, and I appreciate the opportunity to work on this tech assessment. I have about 3 years of software engineering experience, primarily on the backend in Go, with some experience in Python and Ruby on Rails. Prior to that, I have worked in finance, consulting and the military. Let's connect on [LinkedIn](https://www.linkedin.com/in/markcyen/)!
