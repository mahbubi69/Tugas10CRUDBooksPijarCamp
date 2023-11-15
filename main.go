package main

import (
	"fmt"
	"log"
	"os"
	"tugas-10-crud-books-pijar-camp/controller"

	"github.com/joho/godotenv"
)

func main() {
	server := controller.Server{}

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error mengambil env %v", err)
	} else {
		fmt.Println("Succes get data env")
	}

	server.InitializedConnectToDb(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	server.RunServer(":" + os.Getenv("PORT"))
}
