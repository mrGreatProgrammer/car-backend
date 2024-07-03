package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
}

func LoadConfig() Config {
	return Config{
		DBName:     "car_db",
		DBUser:     "postgres",
		DBPassword: "bezhan2009",
		DBHost:     "localhost",
		DBPort:     "5432",
	}
}

func InitDatabase() {
	config := LoadConfig()

	// Формирование строки подключения к PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tashkent",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)

	// Подключение к базе данных
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Применение миграций
	err = DB.AutoMigrate(&Car{}, &User{})
	if err != nil {
		fmt.Println("failed to migrate database Error:", err)
	}
}
