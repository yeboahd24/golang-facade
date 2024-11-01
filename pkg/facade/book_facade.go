package facade

import (
	"github.com/gin-gonic/gin"
	"github.com/yeboahd24/facade/pkg/dto"
	"github.com/yeboahd24/facade/pkg/service"
	"strconv"
)

type BookFacade interface {
	HandleCreateBook(c *gin.Context)
	HandleGetAllBooks(c *gin.Context)
	HandleGetBook(c *gin.Context)
	HandleUpdateBook(c *gin.Context)
	HandleDeleteBook(c *gin.Context)
}

type bookFacade struct {
	service service.BookService
}

func NewBookFacade(service service.BookService) BookFacade {
	return &bookFacade{service: service}
}

func (f *bookFacade) HandleCreateBook(c *gin.Context) {
	var req dto.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	book, err := f.service.CreateBook(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, book)
}

func (f *bookFacade) HandleGetAllBooks(c *gin.Context) {
	books, err := f.service.GetAllBooks()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, books)
}

func (f *bookFacade) HandleGetBook(c *gin.Context) {
	id := uint(c.GetUint("id"))
	book, err := f.service.GetBook(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(200, book)
}

// func (f *bookFacade) HandleUpdateBook(c *gin.Context) {
// 	id := uint(c.GetUint("id"))
// 	var req dto.UpdateBookRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}
//
// 	book, err := f.service.UpdateBook(id, &req)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}
//
// 	c.JSON(200, book)
// }

func (f *bookFacade) HandleUpdateBook(c *gin.Context) {
	// Convert string ID to uint
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	var req dto.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	book, err := f.service.UpdateBook(uint(id), &req)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, book)
}

func (f *bookFacade) HandleDeleteBook(c *gin.Context) {
	id := uint(c.GetUint("id"))
	if err := f.service.DeleteBook(id); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Book deleted successfully"})
}
