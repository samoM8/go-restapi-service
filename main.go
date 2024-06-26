package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-restapi-service/config"
	"github.com/go-restapi-service/routes"
	"log"
	"os"
)

func main() {
	//Connect to db
	config.Connect()

	if os.Getenv("ENV") == "docker" {
		gin.SetMode(gin.ReleaseMode)
	}

	//Init routes
	router := gin.Default()

	//Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":8080"))
}
