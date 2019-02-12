package models

type Paginate struct {
	Page int
	PageSize int
	Data interface{}
	Total int
}
