package routes 
import (
  "github.com/gorilla/mux"
  "github.com/sajagsubedi/Bookstore-Management-Api/pkg/controllers"
  )
  
  var RegisterRoutes=func(router *mux.Router){
    router.HandleFunc("/books",controllers.GetAllBooks).Methods("GET")
    router.HandleFunc("/book{bookid}",controllers.GetBookByID).Methods("GET")
    router.HandleFunc("/addbook",controllers.AddBook).Methods("POST")
    router.HandleFunc("/book/{bookid}",controllers.UpdateBook).Methods("PUT")
    router.HandleFunc("/book/{bookid}",controllers.DeleteBook).Methods("DELETE")
  }