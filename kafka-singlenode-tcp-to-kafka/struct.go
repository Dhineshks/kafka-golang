package main

//Info struct
type Info struct {
	Name       string       `json:"name"`
	FatherName string       `json:"fathers_name"`
	Standard   string       `json:"standard"`
	Section    string       `json:"section"`
	Data       []Datastruct `json:"array"`
}

//Datastruct struct
type Datastruct struct {
	Date      string `json:"date"`
	TotalMark string `json:"total_mark"`
	Grade     string `json:"grade"`
}
