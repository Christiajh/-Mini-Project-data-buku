package controllers

import (
	"Mini-Project-data-buku/config"
	"Mini-Project-data-buku/models"
	"net/http"
	

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	books := []models.Book{}
	rows, err := config.DB.Query("SELECT * FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var b models.Book
		rows.Scan(&b.ID, &b.Title, &b.Description, &b.ImageURL, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID, &b.CreatedAt)
		books = append(books, b)
	}
	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var b models.Book
	err := config.DB.QueryRow("SELECT * FROM books WHERE id=$1", id).
		Scan(&b.ID, &b.Title, &b.Description, &b.ImageURL, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID, &b.CreatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, b)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	res, err := config.DB.Exec("DELETE FROM books WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil dihapus"})
}

func CreateBook(c *gin.Context) {
	var b models.Book
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if b.ReleaseYear < 1980 || b.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tahun rilis harus antara 1980-2024"})
		return
	}
	if b.TotalPage < 100 {
		b.Thickness = "tipis"
	} else {
		b.Thickness = "tebal"
	}
	_, err := config.DB.Exec(`
		INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,now())`,
		b.Title, b.Description, b.ImageURL, b.ReleaseYear, b.Price, b.TotalPage, b.Thickness, b.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Buku berhasil ditambahkan"})
}

func GetBooksByCategoryID(c *gin.Context) {
	id := c.Param("id")
	rows, err := config.DB.Query("SELECT * FROM books WHERE category_id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	books := []models.Book{}
	for rows.Next() {
		var b models.Book
		rows.Scan(&b.ID, &b.Title, &b.Description, &b.ImageURL, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID, &b.CreatedAt)
		books = append(books, b)
	}
	c.JSON(http.StatusOK, books)
}