package dto

type RequestPagination struct {
	Filter     string `query:"filter"`
	PageSize   int    `query:"page_size"`
	PageNumber int    `query:"page_number"`
}
