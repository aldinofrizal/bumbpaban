package main

import (
	"github.com/aldinofrizal/bumpaban/models"
	"github.com/aldinofrizal/bumpaban/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.DBConnect()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	routes.SetupRoutes(r)

	r.Run(":8000")
}
