package models

//Response
type HealthCheckResponse struct {
	Message string `json:"message"`
}


type ErrorResponse struct {
	Status   bool     `json:"status"`
	Messages []string `json:"messages"`
}


type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}