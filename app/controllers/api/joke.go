package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"jokeclub/pkg/e"
	"jokeclub/pkg/util"
	"jokeclub/app/models"
	"fmt"
)

func Jokes(c *gin.Context) {
	var jokes models.Joke
	models.DB.First(jokes)
	fmt.Println(jokes)
	c.JSON(http.StatusOK, util.RetJson(e.Success, jokes))
}
