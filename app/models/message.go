package models

import "time"

type MessageSession struct {
	ID          uint       `gorm:"primary_key"json:"id"`
	FromUserId  uint       `json:"from_user_id"`
	ToUserId    uint       `json:"to_user_id"`
	LastMessage string     `json:"last_message"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
}

type MessageMap struct {
	ID               uint       `gorm:"primary_key"json:"id"`
	MessageSessionId uint       `json:"message_session_id"`
	MessageId        uint       `json:"message_id"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"-"`
	DeletedAt        *time.Time `json:"-"`
}

type Message struct {
	ID        uint       `gorm:"primary_key"json:"id"`
	UserId    uint       `json:"user_id"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

func (MessageSession) TableName() string {
	return "message_sessions"
}

func (MessageMap) TableName() string {
	return "message_maps"
}

func (Message) TableName() string {
	return "messages"
}