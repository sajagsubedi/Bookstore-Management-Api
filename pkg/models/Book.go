package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/sajagsubedi/Bookstore-Management-Api/pkg/config"
)
var db *gorm.DB

type Book struct {
  gorm.Model
  Name string `gorm:""json:"name"`
  Author string `json:"author"`
  Publication string `json:"publication"`
}
func init() {
  config.ConnectDb()
  db = config.GetDb()
  db.AutoMigrate(&Book{})
}
func (b *Book) CreateBook() *Book {
  db.NewRecord(b)
  db.Create(&b)
  return b
}
func GetBooks() []Book{
  var Books []Book
  db.Find(&Books)
  return Books
}
func GetBookById(ID int64) (*Book,*gorm.DB){
  var myBook Book
  db=db.Where("ID=?",ID).Find(&myBook)
  return &myBook,db
}
func DeleteBookById(ID int64) Book{
  var myBook Book
  db.Where("ID=?",ID).Delete(&myBook)
  return myBook
}