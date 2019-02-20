package cache

import (
	"github.com/gin-gonic/gin"
	"jokeclub/app/models"
	"jokeclub/pkg/util"
)

const (
	CACHE_ARTICLE = "ARTICLE"
	CACHE_TAG     = "TAG"
)

func UserProfile(c *gin.Context) models.User {
	token := c.Query("token")
	claims, err := util.ParseToken(token)
	if err != nil {
		c.AbortWithStatus(401)
	}

	return models.GetUserProfile(claims.UserId)
}
