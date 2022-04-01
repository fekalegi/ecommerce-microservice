package models

// JSONResponses : json response
type JSONResponses struct {
	Code   int         `json:"code"`
	Status interface{} `json:"status"`
	Data   interface{} `json:"data"`
}

// JSONResponsesPagination : json with pagination
type JSONResponsesPagination struct {
	Code   int         `json:"code"`
	Status interface{} `json:"status"`
	Data   interface{} `json:"data"`
	Meta   interface{} `json:"meta"`
}

// JSONResponsesSwaggerSucceed : json response for swagger
type JSONResponsesSwaggerSucceed struct {
	Code   int    `json:"code" default:"200" example:"200"`
	Status string `json:"status" default:"Succeed" example:"Succeed"`
	Data   string `json:"data" default:"nil"`
}
