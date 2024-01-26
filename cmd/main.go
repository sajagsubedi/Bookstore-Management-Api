package main
import (
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/sajagsubedi/Bookstore-Management-Api/pkg/routes"
)

func main() {
  r:= mux.NewRouter()
  routes.RegisterRoutes(r)
  http.Handle("/", r)
  fmt.Println("Server running in port 8000")
  if err:= http.ListenAndServe(":8000", r); err != nil {
    log.Fatal(err)
  }
}