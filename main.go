package main

import (
	"io"
	"os"

	"github.com/KadirbekSharau/mywealth-backend/config/db"
	"github.com/KadirbekSharau/mywealth-backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	db := db.NewDatabaseConnection()
	server := gin.Default()

	server.Use(
		cors.New(cors.Config{
			AllowOrigins:  []string{"*"},
			AllowMethods:  []string{"*"},
			AllowHeaders:  []string{"*"},
			AllowWildcard: true,
		}),
	)


	routes.InitAuthRoutes(db, server)
	// routes.InitBookRoutes(db, server)
	// routes.InitUserRoutes(db, server)

	server.Run(":8080")
}