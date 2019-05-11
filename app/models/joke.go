package models

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"jokeclub/pkg/util"
	"time"
)

const (
	TextType  = 0
	PicType   = 1
	VideoTYpe = 2
)

type Joke struct {
	ID          int        `gorm:"primary_key"json:"id"`
	UserId      int        `json:"user_id"`
	User        User       `json:"user"`
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
	DeletedAt   *time.Time `json:"-"`
}

func (Joke) TableName() string {
	return "jokes"
}

func JokePaginate(c *gin.Context, page string, pageSize string, maps interface{}, order string) (p Paginate) {

	pageInt, _ := com.StrTo(page).Int()
	pageSizeInt, _ := com.StrTo(pageSize).Int()
	offset := util.GetPageOffset(pageInt, pageSizeInt)

	var jokes []Joke
	DB.Order(order).Preload("User").Where(maps).Offset(offset).Limit(pageSize).Find(&jokes)

	return Paginate{
		CurrentPage: pageInt,
		PerSize:     pageSizeInt,
		Data:        jokes,
		Total:       GetJokePaginateTotal(maps),
		Path:        c.Request.URL.Path,
	}
}

func GetJokePaginateTotal(maps interface{}) (count int) {
	DB.Model(&Joke{}).Where(maps).Count(&count)
	return count
}

func GetJoke(id int) Joke {
	var joke Joke
	DB.Where(Joke{ID: id}).First(&joke)
	return joke
}

func (joke Joke) Up(userId int, cancel bool) {
	var like Like

	if cancel {
		DB.Where(Like{UserId: userId, JokeId: joke.ID}).First(&like)
		DB.Delete(&like)
		joke.UpNum -= 1
	} else {
		DB.FirstOrCreate(&like, Like{JokeId: joke.ID, UserId: userId})
		joke.UpNum += 1
	}

	DB.Save(&joke)

}

func (joke Joke) Down(userId int, cancel bool) {
	var down Down
	if cancel {
		DB.Where(Down{UserId: userId, JokeId: joke.ID}).First(&down)
		DB.Delete(&down)
		joke.DownNum -= 1
	} else {
		DB.FirstOrCreate(&down, Down{JokeId: joke.ID, UserId: userId})
		joke.DownNum += 1
	}

	DB.Save(&joke)
}
