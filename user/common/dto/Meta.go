package dto

type Meta struct {
	TotalPages      int  `json:"total_pages"`
	PageSize        int  `json:"page_size"`
	TotalItem       int  `json:"total_item"`
	HasNextPage     bool `json:"has_next_page"`
	HasPreviousPage bool `json:"has_previous_page"`
}
