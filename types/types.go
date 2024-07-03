package types

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []struct {
		ShortDescription string `json:"shortDescription"`
		Price            string `json:"price"`
	} `json:"items"`
	Total string `json:"total"`
}

type Err struct {
	Error string `json:"error"`
}

type ReceiptId struct {
	Id string `json:"id"`
}

type Points struct {
	Points int `json:"points"`
}
