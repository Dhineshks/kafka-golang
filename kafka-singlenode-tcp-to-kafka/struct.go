package main

//A Barcode reader sends customer data through tcp

//Info struct
type Info struct {
	CustomerName string       `json:"customer_name"`
	Address      string       `json:"address"`
	Phone        string       `json:"phone"`
	CustomerID   string       `json:"id"`
	Data         []Datastruct `json:"array"`
}

//Datastruct struct
type Datastruct struct {
	PurchasedDate string `json:"date"`
	TotalAmount   string `json:"total_amount"`
}
