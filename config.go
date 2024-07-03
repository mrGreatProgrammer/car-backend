package main

type Config struct {
    DBName     string // Название базы данных
    DBUser     string // Пользователь базы данных
    DBPassword string // Пароль пользователя базы данных
    DBHost     string // Хост базы данных
    DBPort     string // Порт базы данных
}

func LoadConfig() Config {
    return Config{
        DBName:     "car_backend_db",
        DBUser:     "postgres",
        DBPassword: "bezhan2009",
        DBHost:     "localhost",
        DBPort:     "5432",
    }
}
