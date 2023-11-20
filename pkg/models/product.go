package models

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type ProductResponse struct {
	StatusCode int      `json:"statusCode"`
	Product    *Product `json:"product"`
}

type ProductListResponse struct {
	StatusCode int       `json:"statusCode"`
	Products   []Product `json:"products"`
}

type CreateProductRequest struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type SortBy string

const (
	Name     SortBy = "name"
	Quantity SortBy = "quantity"
)

type Order string

const (
	Asc  Order = "asc"
	Desc Order = "desc"
)
