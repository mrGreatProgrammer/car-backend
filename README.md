
# Car

![GitHub top language](https://img.shields.io/github/languages/top/mrGreatProgrammer/car-backend) 
![GitHub language count](https://img.shields.io/github/languages/count/mrGreatProgrammer/car-backend)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/mrGreatProgrammer/car-backend)
![GitHub repo size](https://img.shields.io/github/repo-size/mrGreatProgrammer/car-backend) 
![GitHub last commit](https://img.shields.io/github/last-commit/mrGreatProgrammer/car-backend)
![GitHub User's stars](https://img.shields.io/github/stars/mrGreatProgrammer?style=social)

Read in [Русский](README.ru.md)

The "Car" project is designed for listing cars for sale. Users can add information about their cars they want to sell, and other users can browse these listings.

## Technologies

### Backend

- **Programming Language**: Go
- **Web Framework**: Gin
- **ORM**: GORM
- **Database**: PostgreSQL

## Installation and Launch

### Step 1: Clone the repository

```sh
git clone https://github.com/bezhan2009/car.git
cd car
```

### Step 2: Configure settings

Create a file named `config.go` and add the following settings:

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

### Step 3: Install dependencies

Install the necessary dependencies using the command:

```sh
go mod tidy
```

### Step 4: Launch the application

Run the application:

```sh
go run main.go
```

The application will be available at [http://localhost:8080](http://localhost:8080).

## Main Features

- Add a new car listing for sale
- View the list of car listings
- Search car listings by various criteria

### Contacts

If you have any questions or suggestions about the project, please contact me via email at [karimovbezan0@gmail.com](mailto:karimovbezhan0@gmail.com).

---

Thank you for using the "Car" project! I hope it will be useful to you.
