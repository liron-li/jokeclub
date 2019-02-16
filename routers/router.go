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

	gin.SetMode(setting.RunMode)

	r.GET("/", home.Index)

	r.POST("/api/login", api.Login)

	r.Use(middleware.Cors())

	authRoute := r.Group("/api/user")
	authRoute.Use(middleware.JWT())
	{
		authRoute.GET("/profile", api.Profile)
	}

	jokeRoute := r.Group("/api/jokes")

	jokeRoute.Use()
	{
		jokeRoute.GET("", api.Jokes)
	}

	return r
}
