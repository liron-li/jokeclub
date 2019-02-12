package models

import (
	"jokeclub/pkg/util"
	"github.com/Unknwon/com"
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

func JokePaginate(page string, pageSize string, maps interface{}) (p Paginate) {

	pageInt, _ := com.StrTo(page).Int()
	pageSizeInt, _ := com.StrTo(pageSize).Int()

	var jokes []Joke
	DB.Where(maps).Offset(util.GetPageOffset(pageInt, pageSizeInt)).Limit(pageSize).Find(&jokes)
	return Paginate{Page: pageInt, PageSize: pageSizeInt, Data: jokes, Total: GetJokePaginateTotal(maps)}
}

func GetJokePaginateTotal(maps interface{}) (count int) {
	DB.Model(&Joke{}).Where(maps).Count(&count)
	return count
}
