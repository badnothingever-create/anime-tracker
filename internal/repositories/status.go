package repositories

import (
	"anime-tracker/internal/database"
	//"log"
)

func SaveUserAnimeStatus(userID int, animeID int, status string) error {
	_, err := database.DB.Exec(`
    	INSERT INTO user_anime_status (user_id, anime_id, status)
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id, anime_id)
		DO UPDATE SET status = EXCLUDED.status;
	`, userID, animeID, status)
	if err != nil {
		//log.Printf("Ошибка выполнения SQL для пользователя %d, animeID %d, статус %q: %v", userID, animeID, status, err)
	} else {
		//log.Printf("Успешно сохранено для userID=%d, animeID=%d, статус=%q", userID, animeID, status)
	}
	//log.Printf("Попытка сохранить статус. userID=%d, animeID=%d, status=%q", userID, animeID, status)
	return err
}
