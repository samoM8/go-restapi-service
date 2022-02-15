package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
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

	//Serve swagger.yaml file
	router.GET("/swagger.yaml", gin.WrapH(http.FileServer(http.Dir("./"))))

	//Swagger UI
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	router.GET("/docs", gin.WrapH(sh))
}

func welcome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Welcome to API",
	})
	return
}
