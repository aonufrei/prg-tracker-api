package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"prg-tracker/data"
)

func main() {
	dbConfig := &data.DatabaseConfig{
		Filename: "prg-tracker.db",
	}
	db, err := data.InitDatabase(dbConfig)
	if err != nil {
		log.Fatal("Failed to initialize database connection")
	}
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "I am alive!!!!",
		})
	})

	dependencies, _ := InitDependencies(db)
	InitRoutes(r, dependencies)

	err = r.Run()
	if err != nil {
		log.Fatal("Failed to run application", err)
	}
}
