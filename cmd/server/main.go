package main

import (
	"go-node/internal/database"
	"go-node/internal/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация базы данных
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Инициализация роутера
	r := gin.Default()

	// Статические файлы

	r.Static("/static", "./web/static")

	// HTML шаблоны
	r.LoadHTMLGlob("web/templates/*")

	// Главная страница
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// API маршруты
	api := r.Group("/api")
	{
		api.GET("/nodes", handlers.GetNodes(db))
		api.POST("/nodes", handlers.CreateNode(db))
		api.PUT("/nodes/:id", handlers.UpdateNode(db))
		api.DELETE("/nodes/:id", handlers.DeleteNode(db))
	}

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}