
# Car

![GitHub top language](https://img.shields.io/github/languages/top/mrGreatProgrammer/car-backend) 
![GitHub language count](https://img.shields.io/github/languages/count/mrGreatProgrammer/car-backend)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/mrGreatProgrammer/car-backend)
![GitHub repo size](https://img.shields.io/github/repo-size/mrGreatProgrammer/car-backend) 
![GitHub last commit](https://img.shields.io/github/last-commit/mrGreatProgrammer/car-backend)
![GitHub User's stars](https://img.shields.io/github/stars/mrGreatProgrammer?style=social)

Проект "Car" предназначен для объявления автомобилей на продажу. Пользователи могут добавлять информацию о своих автомобилях, которые они хотят продать, а другие пользователи могут просматривать эти объявления.

## Технологии

### Бэкенд

- **Язык программирования**: Go
- **Веб-фреймворк**: Gin
- **ORM**: GORM
- **СУБД**: PostgreSQL

## Установка и запуск

### Шаг 1: Клонирование репозитория

```sh
git clone https://github.com/bezhan2009/car.git
cd car
```

### Шаг 2: Настройка конфигурации

Создайте файл `config.go` и добавьте в него следующие настройки:

```go
package main

type Config struct {
    DBName     string
    DBUser     string
    DBPassword string
    DBHost     string
    DBPort     string
}

func LoadConfig() Config {
    return Config{
        DBName:     "mydatabase",
        DBUser:     "postgres",
        DBPassword: "password",
        DBHost:     "localhost",
        DBPort:     "5432",
    }
}
```

### Шаг 3: Установка зависимостей

Установите необходимые зависимости с помощью команды:

```sh
go mod tidy
```

### Шаг 4: Запуск приложения

Запустите приложение:

```sh
go run main.go
```

Приложение будет доступно по адресу [http://localhost:8080](http://localhost:8080).

## Основные функции

- Добавление нового объявления о продаже автомобиля
- Просмотр списка объявлений
- Поиск объявлений по различным критериям

### Контакты

Если у вас есть вопросы или предложения по проекту, пожалуйста, свяжитесь со мной по электронной почте [karimovbezhan0@gmail.com](mailto:karimovbezhan0@gmail.com).

---

Спасибо за использование проекта "Car"! Надеюсь, он будет полезен для вас.
