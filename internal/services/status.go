package services

import (
	"anime-tracker/internal/database"
	//"anime-tracker/internal/models"
	"anime-tracker/internal/repositories"
	"errors"
	//"log"
)

var allowedStatuses = map[string]bool{
	"-":                   true,
	"Просмотрено":         true,
	"Смотрю":              true,
	"Планирую":            true,
	"Никогда не посмотрю": true,
}

func UpdateAnimeStatus(userID int, animeID int, status string) error {
	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM animes WHERE id=$1)", animeID).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("аниме с таким ID не найдено")
	}
	// Вызываем функцию репозитория для сохранения
	return repositories.SaveUserAnimeStatus(userID, animeID, status)
}
