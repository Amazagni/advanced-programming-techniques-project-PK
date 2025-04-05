package models

type Item struct {
	Id          int32
	Name        string
	Description string
	Quantity    int32
	ImageURL    string
}

type ErrorResponse struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}
