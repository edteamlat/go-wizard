package model

// Response struct of an error or message
type Response struct {
	Code    StatusCode `json:"code,omitempty"`
	Message string     `json:"message,omitempty"`
}

// Responses slice of Response
type Responses []Response

// MessageResponse contains the response message
type MessageResponse struct {
	Data     interface{} `json:"data,omitempty"`
	Errors   Responses   `json:"errors,omitempty"`
	Messages Responses   `json:"messages,omitempty"`
}
