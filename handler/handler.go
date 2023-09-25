package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sraynitjsr/db"
)

func GetNodes(graphDB db.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		nodes, err := graphDB.GetNodes()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, nodes)
	}
}

func GetNodeByID(graphDB db.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		node, err := graphDB.GetNodeByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if node == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Node not found"})
			return
		}

		c.JSON(http.StatusOK, node)
	}
}

func CreateNode(graphDB db.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		var nodeData struct {
			Name string `json:"name"`
		}

		if err := c.ShouldBindJSON(&nodeData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := graphDB.CreateNode(nodeData.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

func UpdateNode(graphDB db.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var nodeData struct {
			Name string `json:"name"`
		}

		if err := c.ShouldBindJSON(&nodeData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = graphDB.UpdateNode(id, nodeData.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Node updated successfully"})
	}
}

func DeleteNode(graphDB db.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		err = graphDB.DeleteNode(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Node deleted successfully"})
	}
}
