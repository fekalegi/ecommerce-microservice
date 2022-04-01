package dto

type RequestNote struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ResponseBaseNote struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Creator     string `json:"creator"`
}

type ResponseFetchNote struct {
	Note ResponseBaseNote `json:"note"`
}

type ResponseFetchAllNote struct {
	Notes []ResponseBaseNote `json:"notes"`
}

type ResponseCreateNote struct {
	ID int `json:"id"`
}
