package main
import (
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/sajagsubedi/Bookstore-Management-Api/pkg/routes"
)
func main(){
  r:=mux.NewRouter()
  routes.RegisterRoutes(r)
  http.Handle("/",r)
  fmt.Println("Server running in port 8080")
  if err:=http.ListenAndServe(":8080",r);err!=nil{
    log.Fatal(err)
  }
}