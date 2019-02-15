package models

type Paginate struct {
	CurrentPage int         `json:"current_page"`
	PerSize     int         `json:"per_size"`
	Data        interface{} `json:"data"`
	Total       int         `json:"total"`
	Path        string      `json:"path"`
}
