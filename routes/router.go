package routes

import (
	"Mini-Project-data-buku/controllers"
	"Mini-Project-data-buku/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Tambahkan root endpoint agar tidak 404 saat akses domain langsung
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "🚀 API Mini-Project Data Buku is running!",
		})
	})

	api := r.Group("/api")
	api.POST("/users/login", controllers.Login)

	api.GET("/categories", middlewares.JWTAuth(), controllers.GetCategories)
	api.POST("/categories", middlewares.JWTAuth(), controllers.CreateCategory)
	api.GET("/categories/:id", middlewares.JWTAuth(), controllers.GetCategoryByID)
	api.DELETE("/categories/:id", middlewares.JWTAuth(), controllers.DeleteCategory)
	api.GET("/categories/:id/books", middlewares.JWTAuth(), controllers.GetBooksByCategoryID)

	api.GET("/books", middlewares.JWTAuth(), controllers.GetBooks)
	api.POST("/books", middlewares.JWTAuth(), controllers.CreateBook)
	api.GET("/books/:id", middlewares.JWTAuth(), controllers.GetBookByID)
	api.DELETE("/books/:id", middlewares.JWTAuth(), controllers.DeleteBook)
}
