package main

import (
	"github.com/gin-gonic/gin"
	"github.com/satyamksharma/REST-Api-for-Event-Management-with-Authentication.git/db"
	"github.com/satyamksharma/REST-Api-for-Event-Management-with-Authentication.git/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}

