package services

import (
	"anime-tracker/internal/repositories"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

// Регистрация пользователя с хешированием пароля
func RegisterUser(username, password string) error {
	exists, err := repositories.UserExists(username)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("пользователь уже существует")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return repositories.CreateUser(username, string(hash))
}

// Аутентификация пользователя: проверка пароля
func AuthenticateUser(username, password string) (int, error) {
	user, err := repositories.GetUserByUsername(username)
	if err != nil {
		return 0, errors.New("пользователь не найден")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return 0, errors.New("неверные учетные данные")
	}
	return user.ID, nil
}

var (
	sessions = map[string]int{}
	mu       sync.Mutex
)

func CreateSession(userID int) string {
	mu.Lock()
	defer mu.Unlock()
	token := fmt.Sprintf("session-%d", userID) // пример генерации токена, можно улучшить
	sessions[token] = userID
	return token
}

// Проверяет, существует ли сессия с таким токеном
func IsSessionValid(token string) bool {
	mu.Lock()
	defer mu.Unlock()
	_, exists := sessions[token]
	return exists
}

// Функция для получения userID по http.Request, возвращает ошибку если сессии нет или некорректная
func GetUserIDFromSession(r *http.Request) (int, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return 0, errors.New("нет cookie сессии")
	}
	if !IsSessionValid(cookie.Value) {
		return 0, errors.New("сессия недействительна")
	}
	userID := GetUserIDBySession(cookie.Value)
	if userID == 0 {
		return 0, errors.New("пользователь не найден по сессии")
	}
	return userID, nil
}

// Получает userID по токену сессии
func GetUserIDBySession(token string) int {
	mu.Lock()
	defer mu.Unlock()
	return sessions[token]
}
