package models

// HTTP is the struc to representate a error json return in api requests.
type HTTP struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
