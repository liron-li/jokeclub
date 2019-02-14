package models

import (
	"jokeclub/pkg/util"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

type Joke struct {
	Model
	UserId      string
	Content     string
	Image       string
	Video       string
	UpNum       int
	DownNum     int
	FavoriteNum int
	CommentNum  int
	Type        int
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
