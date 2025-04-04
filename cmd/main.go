package main

import (
	"nexa/internal/database"
	"nexa/internal/service"
	"nexa/internal/utils"
)

func main() {
	utils.LoadEnv()
	conn := database.ConnectDB()
	service.StartServer(conn)
}
