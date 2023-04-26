package main

import (
	"github.com/aldinofrizal/bumpaban/models"
	"github.com/aldinofrizal/bumpaban/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"token", "access-control-allow-origin", "content-type"}
	r.Use(cors.New(config))

	models.DBConnect()
	// r.Use(gin.Logger())
	r.Use(gin.Recovery())
	routes.SetupRoutes(r)

	r.Run(":8000")
}
