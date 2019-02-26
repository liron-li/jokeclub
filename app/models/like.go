package models

import "time"

type Like struct {
	ID        int       `gorm:"primary_key"json:"id"`
	UserId    int       `json:"user_id"`
	JokeId    int       `json:"joke_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Like) TableName() string {
	return "likeds"
}

type Down struct {
	ID        int       `gorm:"primary_key"json:"id"`
	UserId    int       `json:"user_id"`
	JokeId    int       `json:"joke_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Down) TableName() string {
	return "downs"
}

type Favorite struct {
	ID        int       `gorm:"primary_key"json:"id"`
	UserId    int       `json:"user_id"`
	JokeId    int       `json:"joke_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Favorite) TableName() string {
	return "downs"
}
