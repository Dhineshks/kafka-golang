package main

//Info struct
type Info struct {
	CustomerName string       `json:"customer_name"`
	Locality     string       `json:"locality"`
	Phone        string       `json:"phone"`
	CustomerID   string       `json:"id"`
	Date         string       `json:"date"`
	Data         []Datastruct `json:"array"`
}

//Datastruct struct
type Datastruct struct {
	Timestamp   string `json:"timestamp"`
	TotalAmount string `json:"total_amount"`
}
