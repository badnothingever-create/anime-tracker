package repositories

import (
	"anime-tracker/internal/database"
	"anime-tracker/internal/models"
)

func GetAllAnimes() ([]models.Anime, error) {
	rows, err := database.DB.Query("SELECT id, title, status FROM animes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var animes []models.Anime
	for rows.Next() {
		var a models.Anime
		err = rows.Scan(&a.ID, &a.Title, &a.Status)
		if err != nil {
			return nil, err
		}
		animes = append(animes, a)
	}
	return animes, nil
}

func AddAnime(a models.Anime) error {
	_, err := database.DB.Exec("INSERT INTO animes (title, status) VALUES ($1, $2)", a.Title, a.Status)
	return err
}

func GetUsernameByUserID(userID int) (string, error) {
	var username string
	err := database.DB.QueryRow("SELECT username FROM users WHERE id=$1", userID).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}
