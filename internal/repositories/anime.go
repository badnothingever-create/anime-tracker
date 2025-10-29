package repositories

import (
	"anime-tracker/internal/database"
	"log"
)

type AnimeWithStatus struct {
	ID           int
	Title        string
	StatusString string
}

func GetAnimesForUser(userID int) ([]AnimeWithStatus, error) {
	rows, err := database.DB.Query(`
        SELECT a.id, a.title, COALESCE(uas.status, '') as status
        FROM animes a
        LEFT JOIN user_anime_status uas ON uas.anime_id = a.id AND uas.user_id = $1
    `, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var animes []AnimeWithStatus
	for rows.Next() {
		var anime AnimeWithStatus
		if err := rows.Scan(&anime.ID, &anime.Title, &anime.StatusString); err != nil {
			return nil, err
		}
		animes = append(animes, anime)
	}
	log.Printf("Получено из базы %d записей", len(animes))
	log.Printf("GetAnimesForUser: найдено %d аниме", len(animes))
	return animes, nil
}

func ListAnimesWithStatus(userID int) ([]AnimeWithStatus, error) {
	rows, err := database.DB.Query(`
        SELECT a.id, a.title, COALESCE(uas.status, '') as status
        FROM animes a
        LEFT JOIN user_anime_status uas ON a.id = uas.anime_id AND uas.user_id = $1
        ORDER BY a.id
    `, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []AnimeWithStatus
	for rows.Next() {
		var anime AnimeWithStatus
		if err := rows.Scan(&anime.ID, &anime.Title, &anime.StatusString); err != nil {
			return nil, err
		}
		result = append(result, anime)
	}
	return result, nil
}
