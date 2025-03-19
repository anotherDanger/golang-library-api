package helper_test

type BookResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author,omitempty"`
}

type WebResponse struct {
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Data   *BookResponse `json:"data,omitempty"`
}

type WebResponses struct {
	Code   int             `json:"code"`
	Status string          `json:"status"`
	Data   *[]BookResponse `json:"data,omitempty"`
}
