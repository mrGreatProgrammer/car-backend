package main

import (
	"car-backend/models"
	"car-backend/routes"
	"fmt"
)

func main() {
	// Инициализация базы данных
	models.InitDatabase()

	// Инициализация маршрутов
	router := routes.SetupRouter()

	// Запуск сервера
	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
