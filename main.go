package main

import (
	"log"

	"github.com/carlosm27/apiexercise/handlers"
	"github.com/carlosm27/apiexercise/model"
	"github.com/gin-gonic/gin"
)

func main() {

	db, err := model.SetupDatabase()
	if err != nil {
		log.Println("Failed setting up database")
	}
	db.DB()

	router := gin.Default()
	server := handlers.NewServer(db)
	server.RegisterRouter(router)

	log.Fatal(router.Run(":10000"))
}
