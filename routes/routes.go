package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)

	// Users
	router.GET("users")
	router.GET("users/:userId")
	router.POST("users")
	router.PUT("users/:userId")
	router.DELETE("users/:userId")

	//Groups
	router.GET("groups")
	router.GET("groups/:groupId")
	router.POST("groups")
	router.PUT("groups/:groupId")
	router.DELETE("groups/:groupId")
}

func welcome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Welcome to API",
	})
	return
}
