package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// DB - глобальная переменная для подключения к базе данных
var DB *sql.DB

// InitDB - функция для инициализации подключения к базе данных
func InitDB() {
	connStr := os.Getenv("DB_CONN")
	if connStr == "" {
		connStr = "user=postgres password=0000 dbname=animetracker sslmode=disable"
	}

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Настройки пула соединений
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(time.Hour)

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Ошибка пинга базы данных: %v", err)
	}

	log.Println("Подключение к базе данных установлено успешно")
}

// CloseDB - завершает работу с базой
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
