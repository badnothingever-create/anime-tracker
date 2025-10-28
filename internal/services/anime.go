package services

import (
	"anime-tracker/internal/repositories"
)

// Оберточная функция для получения аниме с статусами
func GetAnimesForUser(userID int) ([]repositories.AnimeWithStatus, error) {
	return repositories.GetAnimesForUser(userID)
}
