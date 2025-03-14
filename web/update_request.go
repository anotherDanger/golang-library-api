package web

type UpdateRequest struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}
