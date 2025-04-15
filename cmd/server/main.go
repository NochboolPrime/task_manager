package main

import (
	"log"

	"github.com/NochboolPrime/task_manager/pkg/config"
	"github.com/NochboolPrime/task_manager/pkg/database"
	"github.com/NochboolPrime/task_manager/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	database.ConnectPostgres()
	database.ConnectRedis()

	router := gin.Default()
	routes.InitRoutes(router)

	if err := router.Run(); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
