package route

import (
	"github.com/gin-gonic/gin"
	"github.com/yeboahd24/facade/pkg/facade"
)

func SetupBookRoutes(r *gin.Engine, facade facade.BookFacade) {
	books := r.Group("/books")
	{
		books.POST("", facade.HandleCreateBook)
		books.GET("", facade.HandleGetAllBooks)
		books.GET("/:id", facade.HandleGetBook)
		books.PUT("/:id", facade.HandleUpdateBook)
		books.DELETE("/:id", facade.HandleDeleteBook)
	}
}
