package web

type CreateRequest struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}
