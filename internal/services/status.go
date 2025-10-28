package services

import (
	"anime-tracker/internal/models"
	"anime-tracker/internal/repositories"
	"errors"
	"log"
)

var allowedStatuses = map[string]bool{
	"-":                   true,
	"Просмотрено":         true,
	"Смотрю":              true,
	"Планирую":            true,
	"Никогда не посмотрю": true,
}

func UpdateAnimeStatus(userID, animeID int, status string) error {
	if !allowedStatuses[status] {
		log.Printf("Пользователь %d попытался установить недопустимый статус: %s", userID, status)
		return errors.New("недопустимый статус")
	}
	nullStatus := models.NewNullString(status)
	err := repositories.SaveUserAnimeStatus(userID, animeID, nullStatus.String)
	if err != nil {
		log.Printf("Ошибка сохранения статуса в репозитории для пользователя %d: %v", userID, err)
		return err
	}
	log.Printf("Статус успешно сохранен для пользователя %d, animeID %d: %s", userID, animeID, status)
	return nil
}
