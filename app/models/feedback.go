package models

import (
	"jokeclub/pkg/logging"
	"time"
)

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

func SendFeedback(userId int, content string) bool {

	feedback := Feedback{UserId: userId, Content: content}
	if err := DB.Create(&feedback).Error; err != nil {
		logging.Error("创建 feedback 失败", err)
		return false
	}
	return true
}
