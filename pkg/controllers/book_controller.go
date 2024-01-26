package controllers

import(
  "fmt"
  "github.com/sajagsubedi/Bookstore-Management-Api/pkg/models"
  "github.com/sajagsubedi/Bookstore-Management-Api/pkg/utils"
  "encoding/json"
  "net/http"
  "strconv"
  "github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Request for api")
  var books []models.Book
  books = models.GetBooks()
  res, _ := json.Marshal(books)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  bookid := vars["bookid"]
  ID,err := strconv.ParseInt(bookid, 0, 0)
  if err != nil {
    panic(err)
  }
  book,_:=models.GetBookById(ID)
  res,_ := json.Marshal(book)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(res)
}

func AddBook(w http.ResponseWriter,r *http.Request) {
  fmt.Println("Add request")
  bookData := &models.Book{}
  utils.ParseBody(r, bookData)
  fmt.Println("R",r)
  fmt.Println("bookdata",bookData)
  b := bookData.CreateBook()
  fmt.Println("book",b)
  res,_ := json.Marshal(b)
  fmt.Println(res)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(res)

}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
  var book models.Book
  vars := mux.Vars(r)
  bookid := vars["bookid"]
  ID,err := strconv.ParseInt(bookid, 0, 0)
  if err != nil {
    panic(err)
  }
  book = models.DeleteBookById(ID)
  res,_ := json.Marshal(book)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(res)
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
  editBook :=&models.Book{}
  utils.ParseBody(r, editBook)
  vars := mux.Vars(r)
  bookid := vars["bookid"]
  ID,err := strconv.ParseInt(bookid, 0, 0)
  if err != nil {
    panic(err)
  }
  bookDetails,db := models.GetBookById(ID)
  if editBook.Name != "" {
    bookDetails.Name = editBook.Name
  }
  if editBook.Author != "" {
    bookDetails.Author = editBook.Author
  }
  if editBook.Publication != "" {
    bookDetails.Publication = editBook.Publication
  }
db.Save(&bookDetails)
res,_ :=json.Marshal(bookDetails)
 w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(res)
  
}