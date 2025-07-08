package models

import "time"

type Favorite struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    string    `json:"userId" gorm:"not null"`
	RecipeID  int       `json:"recipeId" gorm:"not null"`
	Title     string    `json:"title" gorm:"not null"`
	Image     string    `json:"image"`
	CookTime  string    `json:"cookTime"`
	Servings  string    `json:"servings"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
