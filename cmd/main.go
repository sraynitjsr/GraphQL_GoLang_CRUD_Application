package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sraynitjsr/db"
	"github.com/sraynitjsr/handler"
)

func main() {

	var graphDB db.Database

	var err error

	if graphDB, err = db.NewGraphDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
		return
	}

	router := gin.Default()

	router.GET("/nodes", handler.GetNodes(graphDB))
	router.GET("/nodes/:id", handler.GetNodeByID(graphDB))
	router.POST("/nodes", handler.CreateNode(graphDB))
	router.PUT("/nodes/:id", handler.UpdateNode(graphDB))
	router.DELETE("/nodes/:id", handler.DeleteNode(graphDB))

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
