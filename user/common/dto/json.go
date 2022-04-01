package dto

type JsonResponses struct {
	Status interface{} `json:"status"`
	Data   interface{} `json:"data"`
	Code   int         `json:"code"`
}

type JsonResponsesPagination struct {
	Code   int         `json:"code"`
	Status interface{} `json:"status"`
	Data   interface{} `json:"data"`
	Meta   interface{} `json:"meta"`
}
