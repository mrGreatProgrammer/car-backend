package routes

import (
	"car-backend/controllers"
	"car-backend/models"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Передача экземпляра базы данных в контроллеры
	db := models.DB
	router.GET("/users/:id", func(c *gin.Context) {
		controllers.GetUser(c, db)
	})

	// Пример маршрута для автомобилей
	router.GET("/cars", controllers.GetCars)
	router.POST("/cars", controllers.CreateCar)

	return router
}
