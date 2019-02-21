package routers

import (
	"github.com/gin-gonic/gin"

	"jokeclub/app/controllers/api"
	"jokeclub/app/controllers/home"
	"jokeclub/app/middleware"
	"jokeclub/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.LoadHTMLGlob("resources/views/*")

	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/", home.Index)

	r.Use(middleware.Cors())
	r.POST("/api/login", api.Login)
	r.POST("/api/register", api.Register)

	authRoute := r.Group("/api/user")
	authRoute.Use(middleware.JWT())
	{
		authRoute.GET("/profile", api.Profile)
		authRoute.GET("/my-message", api.MyMessages)
		authRoute.POST("/send-message", api.SendMessage)
		authRoute.GET("/my-up-jokes", api.MyUpedJokes)
		authRoute.GET("/my-favorite", api.MyFavorite)
		authRoute.POST("/my-feedback", api.Feedback)
	}

	jokeRoute := r.Group("/api/jokes")

	jokeRoute.Use()
	{
		jokeRoute.GET("", api.Jokes)
		jokeRoute.POST("/up", api.JokeUp)
		jokeRoute.POST("/down", api.JokeDown)
		jokeRoute.POST("/favorite", api.JokeFavorite)
		jokeRoute.GET("/comments", api.Comments)
		jokeRoute.POST("/comments", api.PostComments)
	}

	return r
}
