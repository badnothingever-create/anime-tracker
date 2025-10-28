package services

import (
	"anime-tracker/internal/models"
	"anime-tracker/internal/repositories"
)

// ListAnimes получает список аниме с проверкой ошибок
func ListAnimes() ([]models.Anime, error) {
	return repositories.GetAllAnimes()
}

// CreateAnime создаёт новое аниме
func CreateAnime(a models.Anime) error {
	return repositories.AddAnime(a)
}
