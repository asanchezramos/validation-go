package model

type Response struct {
	Success bool 		`json:"success"`
	Data	interface{} `json:"data"`
}

type AuthResponse struct {
	Success bool 		`json:"success"`
	Token	string 		`json:"token"`
	Data	interface{} `json:"data"`
}

type ResponseMessage struct {
	Success bool 		`json:"success"`
	Data 	interface{} `json:"data,omitempty"`
	Message	string 		`json:"message"`
}

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}