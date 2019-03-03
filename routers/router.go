package routers

import (
	"github.com/gin-gonic/gin"

	"jokeclub/app/controllers/api"
	"jokeclub/app/controllers/home"
	"jokeclub/app/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	//r.LoadHTMLGlob("resources/views/*")

	r.GET("/", home.Index)

	r.Use(middleware.Cors())
	r.POST("/api/login", api.Login)
	r.POST("/api/register", api.Register)

	apiRoute := r.Group("/api")
	apiRoute.Use(middleware.JWT())
	{
		apiRoute.GET("/user/profile", api.Profile)
		apiRoute.GET("/user/messages", api.Messages)
		apiRoute.POST("/user/send-message", api.SendMessage)
		apiRoute.GET("/user/my-up-jokes", api.MyUpedJokes)
		apiRoute.GET("/user/my-favorite", api.MyFavorite)
		apiRoute.POST("/user/my-feedback", api.Feedback)

		apiRoute.GET("/jokes", api.Jokes)
		apiRoute.POST("/joke/up", api.JokeUp)
		apiRoute.POST("/joke/down", api.JokeDown)
		apiRoute.POST("/joke/favorite", api.JokeFavorite)
		apiRoute.GET("/joke/comments", api.Comments)
		apiRoute.POST("/joke/comments", api.PostComments)
	}

	return r
}
