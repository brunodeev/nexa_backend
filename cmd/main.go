package main

import (
	"nexa/internal/database"
	"nexa/internal/service"
	"nexa/internal/utils"
)

func main() {
	utils.LoadEnv()
	database.ConnectDB()
	service.StartServer()
}
