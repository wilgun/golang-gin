package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "u1522677_pustaka:pustaka123@tcp(109.106.253.101:3306)/u1522677_pustakaApi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db Connection is error")
	}

	fmt.Println("Database connection succeed")

	db.AutoMigrate(book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/book/:id", bookHandler.GetBook)
	v1.GET("/books", bookHandler.GetBooks)
	v1.POST("/books", bookHandler.PostBooksHandler)

	router.Run()
}
