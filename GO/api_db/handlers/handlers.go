package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type List struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func GetLists(c *gin.Context) {
	var lists []List
	if err := db.Find(&lists).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lists)
}

func CreateList(c *gin.Context) {
	var input List
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Data inserted successfully"})
}
