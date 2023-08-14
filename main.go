package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func main() {
	//------ Kết nối CSDL PostgreSQL
	// Tạo chuỗi kết nối đến PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		"localhost", "5432", "hieuhoccode", "pass", "erp_db_name")

	// Kết nối đến CSDL PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	// ping db xem có hoạt động không
	sqlDB, err := db.DB()
	if err = sqlDB.Ping(); err != nil {
		log.Fatal("Failed to connect to the database")
	}

	// Khởi tạo route sử dụng Gin
	router := gin.Default()
	// Định nghĩa API GET "/hello" để trả về "Hello, World!"
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// Chạy server API trên http://localhost:8000
	router.Run("localhost:8000")
}
