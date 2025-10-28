package handlers

import (
	"anime-tracker/internal/services"
	"encoding/json"

	//"html/template"
	"log"
	"net/http"
)

func UpdateAnimeStatusHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil || !services.IsSessionValid(cookie.Value) {
		http.Error(w, "неверный логин или пароль", http.StatusUnauthorized)
		return
	}

	log.Printf("Сессия пользователя: %s", cookie.Value)

	userID := services.GetUserIDBySession(cookie.Value)
	var req struct {
		AnimeID int    `json:"anime_id"`
		Status  string `json:"status"`
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Некорректный запрос", http.StatusBadRequest)
		return
	}

	log.Printf("UpdateAnimeStatus вызван с параметрами: userID=%d, animeID=%d, status=%q", userID, req.AnimeID, req.Status)

	err = services.UpdateAnimeStatus(userID, req.AnimeID, req.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Статус аниме успешно обновлен: userID=%d, animeID=%d, status=%s", userID, req.AnimeID, req.Status)

	w.WriteHeader(http.StatusOK)
}

func AnimeListHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := services.GetUserIDFromSession(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	log.Printf("Получен userID: %d", userID)

	animes, err := services.GetAnimesForUser(userID)
	if err != nil {
		http.Error(w, "Ошибка загрузки аниме", http.StatusInternalServerError)
		return
	}

	log.Printf("Возвращаемые данные: %v", animes)

	err = templates.ExecuteTemplate(w, "index.html", map[string]interface{}{
		"Animes": animes,
	})
	if err != nil {
		log.Printf("Ошибка рендера шаблона: %v", err)
		http.Error(w, "Ошибка отображения", http.StatusInternalServerError)
	}
}
