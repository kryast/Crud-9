package main

import (
	"fmt"
	"log"

	"github.com/kryast/Crud-9.git/database"
	"github.com/kryast/Crud-9.git/models"
	"github.com/kryast/Crud-9.git/router"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to Connect DB", err)
	}

	db.AutoMigrate(&models.Article{})

	r := router.SetupRouter(db)
	fmt.Println("Server running on https://localhost:8080")
	r.Run(":8080")
}
