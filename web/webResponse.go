package web

type WebResponse struct {
	Code   int           `json:"code"`
	Status string        `json:"status"`
	Data   *BookResponse `json:"data,omitempty"`
}

type WebResponses struct {
	Code   int             `json:"code"`
	Status string          `json:"status"`
	Data   []*BookResponse `json:"data,omitempty"`
}
