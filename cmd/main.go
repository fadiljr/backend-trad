// main.go
package main

import (
	"log"

	"go-gorm-postgresql/controllers" // Tambahkan ini
	"go-gorm-postgresql/models"
	"go-gorm-postgresql/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() {
	dsn := "host=localhost user=postgres password=postgre dbname=db_gorm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to connect to the database:", err)
	}

	// Set DB global ke controller agar bisa digunakan di handler
	controllers.DB = db

	// Migrasi struktur tabel berdasarkan model
	db.AutoMigrate(&models.User{})
}

func main() {
	// Inisialisasi database
	initDB()

	// Inisialisasi server
	r := gin.Default()

	// Setup routes
	routes.AuthRoutes(r)
	routes.UserRoutes(r)

	// Jalankan server pada port 8080
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
