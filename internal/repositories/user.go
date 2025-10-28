package repositories

import (
	database "anime-tracker/internal/database"
	_ "errors"
)

type User struct {
	ID           int
	Username     string
	PasswordHash string
}

// Создать нового пользователя
func CreateUser(username, passwordHash string) error {
	_, err := database.DB.Exec(`INSERT INTO users (username, password_hash) VALUES ($1, $2)`, username, passwordHash)
	if err != nil {
		return err
	}
	return nil
}

// Проверить, существует ли пользователь с таким именем
func UserExists(username string) (bool, error) {
	var exists bool
	err := database.DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`, username).Scan(&exists)
	return exists, err
}

// Получить пользователя по имени
func GetUserByUsername(username string) (*User, error) {
	var user User
	err := database.DB.QueryRow(`SELECT id, username, password_hash FROM users WHERE username = $1`, username).
		Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
