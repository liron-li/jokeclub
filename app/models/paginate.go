package models

type Paginate struct {
	CurrentPage int
	PerSize     int
	Data        interface{}
	Total       int
	Path        string
}
