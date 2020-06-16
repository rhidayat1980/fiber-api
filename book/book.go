package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"

	// dialect for sql
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rhidayat1980/fiber-api/database"
)

// Book struct
type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

// GetBooks function to get all the books
func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

// GetBook function to get single book
func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

// NewBook function to create a book
func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	var book Book
	book.Title = "1984"
	book.Author = "George Orwell"
	book.Rating = 5
	db.Create(&book)
	c.JSON(book)
}

// DeleteBook function to delete a book
func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No Book Found with ID")
		return
	}
	db.Delete(&book)
	c.Send("Book Successfully deleted")
}
