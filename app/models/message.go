package models

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"jokeclub/pkg/util"
)

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

func MessageSessionPaginate(c *gin.Context, page string, pageSize string, maps interface{}, order string) (p Paginate) {

	pageInt, _ := com.StrTo(page).Int()
	pageSizeInt, _ := com.StrTo(pageSize).Int()
	offset := util.GetPageOffset(pageInt, pageSizeInt)

	var messageSessions []MessageSession
	DB.Order(order).Where(maps).Offset(offset).Limit(pageSize).Find(&messageSessions)

	return Paginate{
		CurrentPage: pageInt,
		PerSize:     pageSizeInt,
		Data:        messageSessions,
		Total:       GetMessageSessionPaginateTotal(maps),
		Path:        c.Request.URL.Path,
	}
}

func GetMessageSessionPaginateTotal(maps interface{}) (count int) {
	DB.Model(&MessageSession{}).Where(maps).Count(&count)
	return count
}