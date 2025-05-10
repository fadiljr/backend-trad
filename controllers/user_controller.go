// controllers/user_controller.go
package controllers

import (
	"go-gorm-postgresql/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

// GetAllUsers mengambil semua data user
func GetAllUsers(c *gin.Context) {
	var users []models.User

	// Ambil query parameter ?page= dan ?limit=
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	// Konversi string ke int
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	offset := (pageInt - 1) * limitInt

	if err := DB.Limit(limitInt).Offset(offset).Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUser mengambil user berdasarkan ID
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser menambahkan user baru
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// UpdateUser memperbarui data user berdasarkan ID
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// DeleteUser menghapus user berdasarkan ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
func GetProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.User
	if err := DB.First(&user, userID.(uint)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Jangan kirim password
	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"age":   user.Age,
	})
}
