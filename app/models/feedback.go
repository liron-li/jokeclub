package models

import "time"

type Feedback struct {
	ID        int       `gorm:"primary_key"json:"id"`
	UserId    int       `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Feedback) TableName() string {
	return "feedbacks"
}
