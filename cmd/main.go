package main

import (
	"fmt"
	"log"
	"nexa/internal/database"
	"nexa/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ocorreu um erro ao carregar o .env: %v", err)
	}

	database.ConnectDB()
	fmt.Println("Deu bom a conexão do DB")
	service.StartServer()
}
