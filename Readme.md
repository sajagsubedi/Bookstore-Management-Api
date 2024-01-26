# Bookstore Management API

This API, written in Golang, serves as a bookstore management system with basic CRUD operations for books. It utilizes Gorm for database interaction, Gorilla Mux for routing, and MySQL as the underlying database.

## Project Structure


- **cmd:** Contains the main application file.
- **pkg:** Houses various subfolders for utilities, configuration, routing, models, and controllers.
- cmd
  - main.go
- pkg
  - utils
  - config
  - routes
  - models
  - controllers
   
## API Endpoints

1. **GET /books:** Retrieve all books.
2. **GET /book/{bookid}:** Get a specific book by ID.
3. **POST /addbook:** Add a new book.
4. **PUT /book/{bookid}:** Update a book by ID.
5. **DELETE /book/{bookid}:** Delete a book by ID.

## Book Information

A book in this system is defined by three attributes:
- Name
- Author
- Publication

## Functions and Flow

- **controllers:** Contains functions for handling HTTP requests and calling respective functions from the models.
  
  - Functions:
    - CreateBook
    - GetBooks
    - GetBookById
    - UpdateBook (utilizes GetBookById and db.Save())
    - DeleteBookById

- **models:** Houses functions related to database operations.

  - Functions:
    - CreateBook
    - GetBooks
    - GetBookById
    - DeleteBookById

## Database Configuration

The database connection details are specified in the `config` folder.

## How to Run

1. Ensure Golang is installed.
2. Navigate to the `cmd` directory.
3. Run `go run main.go`.

## Dependencies

- Gorm
- Gorilla Mux
- MySQL

## Note

- Ensure that the MySQL database is set up and the connection details are correctly configured in the `config` folder.

Feel free to customize and expand upon this bookstore management API based on your specific needs.
