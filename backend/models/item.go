package models

type Item struct {
	Id          int32  `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Quantity    int32  `json:"Quantity"`
	ImageURL    string `json:"ImageURL"`
}

type ErrorResponse struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}
