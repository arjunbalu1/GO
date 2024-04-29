package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type list struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ==================================================================
func main() {
	// Define the connection parameters
	host := "db"
	port := 5432
	user := "postgres"
	password := "password"
	dbname := "mydb"

	// Construct the connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Connect to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Ping the database to check the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to the database!")
	router := gin.Default()
	router.GET("/mydb", func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM list")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var fetchedLists []list

		for rows.Next() {
			var fetchedList list
			err := rows.Scan(&fetchedList.ID, &fetchedList.Name)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			fetchedLists = append(fetchedLists, fetchedList)
		}

		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusOK, fetchedLists)
	})
	router.POST("/mydb", func(c *gin.Context) {
		var input list
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Insert the data into the database
		_, err := db.Exec("INSERT INTO list (id, name) VALUES ($1, $2)", input.ID, input.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Data inserted successfully"})
	})

	router.Run(":8080")
}
