package models

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"jokeclub/pkg/util"
	"time"
)

type Like struct {
	ID        int       `gorm:"primary_key"json:"id"`
	UserId    int       `json:"user_id"`
	JokeId    int       `json:"joke_id"`
	Joke 	  Joke 		`json:"joke"`
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
	Joke 	  Joke 		`json:"joke"`
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
	Joke 	  Joke 		`json:"joke"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Favorite) TableName() string {
	return "favorites"
}

func FavoritePaginate(c *gin.Context, page string, pageSize string, maps interface{}, order string) (p Paginate) {

	pageInt, _ := com.StrTo(page).Int()
	pageSizeInt, _ := com.StrTo(pageSize).Int()
	offset := util.GetPageOffset(pageInt, pageSizeInt)

	var favorites []Favorite
	DB.Order(order).Preload("Joke").Where(maps).Offset(offset).Limit(pageSize).Find(&favorites)

	return Paginate{
		CurrentPage: pageInt,
		PerSize:     pageSizeInt,
		Data:        favorites,
		Total:       favoritePaginateTotal(maps),
		Path:        c.Request.URL.Path,
	}
}

func favoritePaginateTotal(maps interface{}) (count int) {
	DB.Model(&Favorite{}).Where(maps).Count(&count)
	return count
}

func LikeJokesPaginate(c *gin.Context, page string, pageSize string, maps interface{}, order string) (p Paginate) {

	pageInt, _ := com.StrTo(page).Int()
	pageSizeInt, _ := com.StrTo(pageSize).Int()
	offset := util.GetPageOffset(pageInt, pageSizeInt)

	var likes []Like
	DB.Order(order).Preload("Joke").Where(maps).Offset(offset).Limit(pageSize).Find(&likes)

	return Paginate{
		CurrentPage: pageInt,
		PerSize:     pageSizeInt,
		Data:        likes,
		Total:       likeJokesPaginateTotal(maps),
		Path:        c.Request.URL.Path,
	}
}

func likeJokesPaginateTotal(maps interface{}) (count int) {
	DB.Model(&Favorite{}).Where(maps).Count(&count)
	return count
}