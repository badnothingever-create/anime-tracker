package repositories

import (
	"anime-tracker/internal/database"
	"anime-tracker/internal/models"
	//"log"
)

type AnimeWithStatus struct {
	ID           int
	Title        string
	StatusString string
}

func GetAnimesForUser(userID int) ([]AnimeWithStatus, error) {
	//log.Printf("GetAnimesForUser вызывается с userID=%d", userID)
	rows, err := database.DB.Query(`
        SELECT a.id, a.title, uas.status
        FROM animes a
        LEFT JOIN user_anime_status uas ON uas.anime_id = a.id AND uas.user_id = $1
		ORDER BY a.id
    `, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//log.Println("Начинаем загрузку аниме из базы")
	var animes []AnimeWithStatus
	for rows.Next() {
		var a models.Anime
		if err := rows.Scan(&a.ID, &a.Title, &a.Status); err != nil {
			//log.Printf("Ошибка сканирования строки: %v", err)
			return nil, err
		}
		//log.Printf("DB row: Anime ID=%d, Title=%s, Status=%v (Valid=%v)", a.ID, a.Title, a.Status.String, a.Status.Valid)
		animeWithStatus := AnimeWithStatus{
			ID:           a.ID,
			Title:        a.Title,
			StatusString: a.StatusString(),
		}
		animes = append(animes, animeWithStatus)
		if err := rows.Err(); err != nil {
			//log.Printf("Ошибка после итерации по rows: %v", err)
			return nil, err
		}
	}
	//log.Printf("Количество аниме после выборки: %d", len(animes))
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
