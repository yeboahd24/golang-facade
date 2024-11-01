package main

import (
	"github.com/yeboahd24/facade/pkg/config"
	"github.com/yeboahd24/facade/pkg/facade"
	"github.com/yeboahd24/facade/pkg/model"
	"github.com/yeboahd24/facade/pkg/repository"
	"github.com/yeboahd24/facade/pkg/route"
	"github.com/yeboahd24/facade/pkg/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Database configuration
	dbConfig := &config.DBConfig{
		Host:     "localhost",
		User:     "postgres",
		Password: "mesika",
		DBName:   "bookstore",
		Port:     "5432",
	}

	// Initialize DB
	db := config.NewDBConnection(dbConfig)

	// Auto migrate the schema
	if err := db.AutoMigrate(&model.Book{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize layers
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookFacade := facade.NewBookFacade(bookService)

	// Setup Gin router
	r := gin.Default()

	// Setup routes
	route.SetupBookRoutes(r, bookFacade)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
