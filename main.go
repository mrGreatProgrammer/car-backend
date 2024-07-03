package main

import (
    "fmt"
    "my-gin-project/routes"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var db *gorm.DB

func main() {
    config := LoadConfig()

    // Формирование строки подключения к PostgreSQL
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tashkent",
        config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)

    // Подключение к базе данных
    var err error
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Закрытие соединения с базой данных при выходе из приложения
    defer db.Close()

    // Применение миграций (опционально)
    // db.AutoMigrate(&models.User{})

    // Инициализация маршрутов
    router := routes.SetupRouter()

    // Запуск сервера
    err = router.Run(":8080")
    if err != nil {
        fmt.Println(err)
    }
}
