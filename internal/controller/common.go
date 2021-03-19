package controller

type ResponseError struct {
	Message string `json:"message"`
}

type SimpleResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
