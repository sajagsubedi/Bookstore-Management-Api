package main
import (
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "github.com/sajagsubedi/Bookstore-Management-Api/pkg/routes"
)

func main() {
  r:= mux.NewRouter()
  corsHandler := handlers.CORS(
    handlers.AllowedOrigins([]string{"*"}),
    handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH"}),
    handlers.AllowedHeaders([]string{ "Content-Type"}),
  )
 handler:=corsHandler(r)
  
  routes.RegisterRoutes(r)
  http.Handle("/", r)
  fmt.Println("Server running in port 8000")
  if err:= http.ListenAndServe(":8000", handler); err != nil {
    log.Fatal(err)
  }
}