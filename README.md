# api-exercise
 A CRUD API written in Go using Gin and Gorm

 # Requirements
 - Go installed (1.21)
 - SQLite installed

 # Installation
 1. Clone this repo.
 2. Open the console in the project directory.
    ![image](https://github.com/carlosm27/api-exercise/assets/70811425/f6ec4707-53d3-42df-9844-064d9d14fd10)

 3. Run the `go mod tidy` to install all the dependencies.


 4. Run `go run main.go` command to start the server.
    ![image](https://github.com/carlosm27/api-exercise/assets/70811425/736b55d2-322c-4489-925d-ccf061b7de99)


    
 # Use
 The server will start on port 10000. It has five routes:
 - GET `/product` for retrieving all the rows in the database.
 - GET `/product/:sku` for retrieving a product by its sku.
 - POST `/product` for creating a product.
 - PUT `/product/:sku` for updating a product.
 - DELETE `/product/:sku` for deleting a product.

 # Architectural Decision
 I keep the handlers and the database setup separate in different folders to test them and reuse them.
 Also in this case I'm using SQLite, and it will not be difficult to change to other database, without modifying the handlers.
 
 This decision makes the code maintainable if the API grows in more routes and models. Any issue relate to the handlers will be address by looking in the handlers folder.
 The same goes to any issue with the database.
