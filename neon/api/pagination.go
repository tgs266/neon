package api

type PaginationResponse[T any] struct {
	Items []T `json:"items"`
	Total int `json:"total"`
}
