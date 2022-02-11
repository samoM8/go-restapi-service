package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-restapi-service/controllers"
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
	router.GET("groups", controllers.GetAllGroups)
	router.GET("groups/:groupId", controllers.GetSingleGroup)
	router.POST("groups", controllers.CreateGroup)
	router.PUT("groups/:groupId", controllers.EditGroup)
	router.DELETE("groups/:groupId", controllers.DeleteGroup)
}

func welcome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Welcome to API",
	})
	return
}
