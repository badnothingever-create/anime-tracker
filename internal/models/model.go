package models

import (
	"database/sql"
	//"log"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"` // не возвращайте пароль клиенту
}

type Anime struct {
	ID     int            `json:"id"`
	Title  string         `json:"title"`
	Status sql.NullString `json:"status"` // пример: просмотрено, смотрю, планирую посмотреть
}

// Метод возвращает строку статуса или пустую строку, если статус null
func (a Anime) StatusString() string {
	if a.Status.Valid {
		//log.Printf("Anime ID %d: статус = %s", a.ID, a.Status.String)
		return a.Status.String
	}
	return ""
}

type UserAnimeStatus struct {
	AnimeID int    `json:"anime_id"`
	Status  string `json:"status"`
}

func NewNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
