// main.go

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"./api/handlers" // замените "yourproject" на путь к вашему пакету обработчиков

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
    // Загрузка переменных окружения из файла .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Получение данных для подключения к базе данных из переменных окружения
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    // Формирование строки подключения
    dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

    // Подключение к базе данных
    db, err := sql.Open("postgres", dbURI)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    r := mux.NewRouter()

    // User routes
    r.HandleFunc("/api/v2/add", handlers.CreateUser(db)).Methods("POST")

    // Start server
    log.Println("Server is running on port 8000")
    http.ListenAndServe(":8000", r)
}
