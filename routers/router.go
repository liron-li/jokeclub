package routers

import (
	"github.com/gin-gonic/gin"

	"goweb/app/controllers/api"
	"goweb/app/controllers/home"
	"goweb/app/middleware"
	"goweb/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.LoadHTMLGlob("resources/views/*")

	gin.SetMode(setting.RunMode)

	r.GET("/", home.Index)

	r.POST("/api/login", api.Login)

	apiRoute := r.Group("/api")
	apiRoute.Use(middleware.Example(), middleware.JWT())
	{
		apiRoute.GET("/user/profile", api.Profile)
	}

	return r
}
