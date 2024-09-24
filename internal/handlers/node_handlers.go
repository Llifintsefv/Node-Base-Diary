package handlers

import (
	"fmt"
	"go-node/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetNodes(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var nodes []models.Node
		if err := db.Find(&nodes).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		db.Find(&nodes)
		for _, node := range nodes {
			fmt.Println(node.X,node.Y)
		}
		c.JSON(http.StatusOK, nodes)
	}
}

func CreateNode(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var node models.Node
		if err := c.ShouldBindJSON(&node); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&node).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, node)
	}
}

func UpdateNode(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var node models.Node
		if err := db.First(&node, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Node not found"})
			return
		}
		if err := c.ShouldBindJSON(&node); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Save(&node)
		c.JSON(http.StatusOK, node)
	}
}

func DeleteNode(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var node models.Node
		if err := db.First(&node, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Node not found"})
			return
		}
		db.Delete(&node)
		c.JSON(http.StatusOK, gin.H{"message": "Node deleted successfully"})
	}
}