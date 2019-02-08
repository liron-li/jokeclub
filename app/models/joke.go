package models

type Joke struct {
	Model
	UserId string
	Content string
}

func (Joke) TableName() string  {
	return "jokes"
}