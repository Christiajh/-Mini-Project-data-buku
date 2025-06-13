package controllers

import (
	"Mini-Project-data-buku/config"
	"Mini-Project-data-buku/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	categories := []models.Category{}
	rows, err := config.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var cat models.Category
		rows.Scan(&cat.ID, &cat.Name)
		categories = append(categories, cat)
	}
	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	var cat models.Category
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := config.DB.Exec("INSERT INTO categories (name, created_at) VALUES ($1, now())", cat.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Kategori berhasil ditambahkan"})
}

func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	var cat models.Category
	err := config.DB.QueryRow("SELECT id, name FROM categories WHERE id=$1", id).
		Scan(&cat.ID, &cat.Name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, cat)
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	res, err := config.DB.Exec("DELETE FROM categories WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Kategori berhasil dihapus"})
}