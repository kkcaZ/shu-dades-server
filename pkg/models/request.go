package models

type Request struct {
	// The request route e.g. /api/v1/users
	Route string `json:"route"`

	// The request Type, e.g. GET, POST, PUT, DELETE
	Type string `json:"type"`

	// The request body
	Body interface{} `json:"body"`

	// The request headers
	Headers string `json:"headers"`
}
