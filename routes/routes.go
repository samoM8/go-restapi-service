package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-restapi-service/controllers"
	"net/http"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)

	// Users
	router.GET("users", controllers.GetAllUsers)
	router.GET("users/:userId", controllers.GetSingleUser)
	router.POST("users", controllers.CreateUser)
	router.PUT("users/:userId", controllers.EditUser)
	router.DELETE("users/:userId", controllers.DeleteUser)

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
