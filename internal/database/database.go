package database

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() *pgx.Conn {
	DB_URI := os.Getenv("DB_URI")

	connConfig, err := pgx.ParseConfig(DB_URI)
	if err != nil {
		log.Fatalf("failed to parse DB URI: %v", err)
	}

	connConfig.DialFunc = func(ctx context.Context, network, addr string) (net.Conn, error) {
		return net.Dial("tcp4", addr)
	}

	conn, err := pgx.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	fmt.Println("Database connection successful.")
	return conn
}
