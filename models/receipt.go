package models

type Receipts struct {
	Receipts []Receipt `json:"receipts"`
}

type Receipt struct {
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Items        []Item  `json:"items"`
	Total        float64 `json:"total"`
}

type Item struct {
	ShortDescription string  `json:"shortDescription"`
	Price       float64 `json:"price"`
}
