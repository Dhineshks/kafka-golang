package main

//Info struct
type Info struct {
	Namespace     string       `json:"namespace"`
	FleetID       string       `json:"fleet_id"`
	DriverID      string       `json:"driver_id"`
	BusinessGroup string       `json:"business_group"`
	Date          string       `json:"date"`
	Data          []Datastruct `json:"array"`
}

//Datastruct struct
type Datastruct struct {
	Timestamp int     `json:"timestamp"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	Speed     float64 `json:"speed"`
	Distance  float64 `json:"distance"`
	EveID     int     `json:"eve_id"`
	EveVal    float64 `json:"eve_val"`
}
