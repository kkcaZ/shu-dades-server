package models

type Request struct {
	// The request route e.g. /api/v1/users
	Route string `json:"route"`

	// The request Type, e.g. GET, POST, PUT, DELETE
	Type RequestType `json:"type"`

	// The request body
	Body interface{} `json:"body"`

	// The request headers
	Headers map[string]string `json:"headers"`
}

type RequestType string

const (
	GET    RequestType = "GET"
	POST   RequestType = "POST"
	PUT    RequestType = "PUT"
	DELETE RequestType = "DELETE"
)

type RequestById struct {
	Id string `json:"id"`
}

type SearchRequest struct {
	PageNumber int    `json:"pageNumber"`
	PageSize   int    `json:"pageSize"`
	SortBy     SortBy `json:"sortBy"`
	Order      Order  `json:"order"`
}
