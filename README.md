# fetch-challenge

directory structure:
├── go.mod                          
├── go.sum                          
├── main.go                         # Entry point for endpoints
├── handlers/
│   ├── receipt_process_handler.go  # Handles receipt processing (POST /receipts/process)
│   ├── points_handler.go           # Handles points retrieval (GET /receipts/:id/points)
├── models/
│   └── receipt.go                  # Defines the Receipt and related models
├── services/
│   └── points_service.go           # Business logic for calculating points
└── utils/
    └── id_generator.go             # Utility for generating unique IDs

