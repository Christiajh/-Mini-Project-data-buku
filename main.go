package main

import (
	"Mini-Project-data-buku/config"
	"Mini-Project-data-buku/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}