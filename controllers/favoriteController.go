package controllers

import (
	"net/http"
	"strconv"

	"github.com/GitNinja36/backend/config"
	"github.com/GitNinja36/backend/models"
	"github.com/gin-gonic/gin"
)

// check the health of the
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// Add to Favorite
func AddFavorite(c *gin.Context) {
	var favorite models.Favorite
	if err := c.ShouldBindJSON(&favorite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if favorite.UserID == "" || favorite.RecipeID == 0 || favorite.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	config.DB.Create(&favorite)
	c.JSON(http.StatusCreated, favorite)
}

// Get Favorites
func GetFavorites(c *gin.Context) {
	userId := c.Param("userId")

	var favorites []models.Favorite
	config.DB.Where("user_id = ?", userId).Find(&favorites)
	c.JSON(http.StatusOK, favorites)
}

// Delete Favorite
func DeleteFavorite(c *gin.Context) {
	userId := c.Param("userId")
	recipeIdStr := c.Param("recipeId")
	recipeId, err := strconv.Atoi(recipeIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
		return
	}

	config.DB.Where("user_id = ? AND recipe_id = ?", userId, recipeId).Delete(&models.Favorite{})
	c.JSON(http.StatusOK, gin.H{"message": "Favorite removed successfully"})
}
