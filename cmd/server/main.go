package main

import (
	database "anime-tracker/internal/database"
	"anime-tracker/internal/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Ошибка открытия файла логов: %v", err)
	}
	defer file.Close()
	log.SetOutput(file)

	log.Println("Запуск сервера")
	// Инициализируем базу данных
	database.InitDB()

	// Если хотите корректный порт без двоеточия
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Создаем новый mux
	mux := http.NewServeMux()

	// Регистрируем роуты, передавая mux
	handlers.InitRoutes(mux)

	log.Printf("Сервер запущен на порту %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
