package main

import (
	"fmt"
	"log"
	"login/config/connection"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("log public path" + pwd)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := connection.App{}
	app.Initialize(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"))

	app.RunServer()
}
