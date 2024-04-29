package main

import (
	"api_db/handlers"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// type List struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// }

var db *gorm.DB

func main() {
	dsn := "host=db user=postgres password=password dbname=mydb port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	handlers.SetDB(db) // Set the database for handlers
	router := gin.Default()

	// Routes
	router.GET("/mydb", handlers.GetLists)
	router.POST("/mydb", handlers.CreateList)

	fmt.Println("Server running at :8080")
	router.Run(":8080")
}

// func getLists(c *gin.Context) {
// 	var lists []List
// 	if err := db.Find(&lists).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, lists)
// }

// func createList(c *gin.Context) {
// 	var input List
// 	if err := c.BindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	if err := db.Create(&input).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, gin.H{"message": "Data inserted successfully"})
// }
