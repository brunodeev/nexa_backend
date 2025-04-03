package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectDB() {
	err := godotenv.Load("internal/database/.env")
	if err != nil {
		log.Println("Aviso: Nenhum arquivo .env foi carregado, usando variáveis do sistema")
	}

	DB_URI := os.Getenv("DB_URI")
	if DB_URI == "" {
		log.Fatal("Erro: A variável de ambiente DB_URI não está definida")
	}

	DB, err = pgx.Connect(context.Background(), DB_URI)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	log.Println("Banco de dados conectado com sucesso!")
}
