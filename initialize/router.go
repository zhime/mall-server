package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/info", getUserInfo)
	}

	return r
}

func getUserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "user info"})
}
