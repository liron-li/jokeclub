package models

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"goweb/pkg/logging"
	"jokeclub/pkg/util"
	"time"
)

type MessageSession struct {
	ID          int        `gorm:"primary_key"json:"id"`
	FromUserId  int        `json:"from_user_id"`
	ToUserId    int        `json:"to_user_id"`
	LastMessage string     `json:"last_message"`
	IsRead      int        `json:"is_read"gorm:"default:'0'"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
}

type MessageMap struct {
	ID               int        `gorm:"primary_key"json:"id"`
	MessageSessionId int        `json:"message_session_id"`
	MessageId        int        `json:"message_id"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"-"`
	DeletedAt        *time.Time `json:"-"`
}

type Message struct {
	ID        int        `gorm:"primary_key"json:"id"`
	UserId    int        `json:"user_id"`
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

func SendMessage(sessionId int, fromUserId int, toUserId int, content string) bool {

	var messageSession MessageSession

	if sessionId <= 0 { // 新增会话
		messageSession = MessageSession{FromUserId: fromUserId, ToUserId: toUserId, LastMessage: content}
		if err := DB.Create(&messageSession).Error; err != nil {
			logging.Error("创建 messageSession 失败", err)
			return false
		}

	} else { // 回复会话

		DB.Where(MessageSession{ID: sessionId}).First(&messageSession)

		if fromUserId != messageSession.FromUserId && fromUserId != messageSession.ToUserId {
			logging.Info("伪造私信", sessionId, fromUserId, toUserId, content)
			return false
		}

		messageSession.IsRead = 0
		messageSession.LastMessage = content

		DB.Save(&messageSession)
	}

	message := Message{UserId: fromUserId, Content: content}
	if err := DB.Create(&message).Error; err != nil {
		logging.Error("创建 message 失败", err)
		return false
	}

	messageMap := MessageMap{MessageSessionId: messageSession.ID, MessageId: message.ID}
	if err := DB.Create(&messageMap).Error; err != nil {
		logging.Error("创建 messageMap 失败", err)
		return false
	}

	return true
}
