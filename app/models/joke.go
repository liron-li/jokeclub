package models

import (
	"jokeclub/pkg/util"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"time"
)

type Joke struct {
	ID          uint       `gorm:"primary_key"json:"id"`
	UserId      string     `json:"user_id"`
	Content     string     `json:"content"`
	Image       string     `json:"image"`
	Video       string     `json:"video"`
	UpNum       int        `json:"up_num"`
	DownNum     int        `json:"down_num"`
	FavoriteNum int        `json:"favorite_num"`
	CommentNum  int        `json:"comment_num"`
	Type        int        `json:"type"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index"json:"deleted_at"`
}

func (Joke) TableName() string {
	return "jokes"
}

func JokePaginate(c *gin.Context, page string, pageSize string, maps interface{}) (p Paginate) {

	pageInt, _ := com.StrTo(page).Int()
	pageSizeInt, _ := com.StrTo(pageSize).Int()

	var jokes []Joke
	DB.Where(maps).Offset(util.GetPageOffset(pageInt, pageSizeInt)).Limit(pageSize).Find(&jokes)
	return Paginate{CurrentPage: pageInt, PerSize: pageSizeInt, Data: jokes, Total: GetJokePaginateTotal(maps), Path: c.Request.URL.Path}
}

func GetJokePaginateTotal(maps interface{}) (count int) {
	DB.Model(&Joke{}).Where(maps).Count(&count)
	return count
}
