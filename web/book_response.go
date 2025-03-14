package web

type BookResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}
