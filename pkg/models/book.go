package models

import (
	"github.com/dibyanshu-mohanty/my-book-sql/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// -> Routes -> Controllers (for Fucntionality) -> Models (for DB)

func (b *Book) CreateBook() *Book {
	db.Attrs(b)
	db.Create(&b)
	return b
}

func GeatAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(id int64) Book {
	var getBook Book
	db.Where("ID=?", id).Delete(&getBook)
	return getBook
}
