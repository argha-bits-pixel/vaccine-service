package main

import (
	"log"
	"vaccination-service/adapters/mysql"
	"vaccination-service/server"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading env file", err.Error())
	}
}

func main() {
	server.Start()
	mysql.Close()
}
