package handlers

import (
	"anime-tracker/internal/services"
	"bytes"
	"encoding/json"
	"io"
	"strconv"
	"strings"

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

	//log.Printf("Статус аниме успешно обновлен: userID=%d, animeID=%s, status=%s", userID, req.AnimeID, req.Status)

	bodyBytes, _ := io.ReadAll(r.Body)
	log.Printf("Request body: %s", string(bodyBytes))
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // чтобы потом можно было прочесть тело снова

	log.Printf("Сессия пользователя: %s", cookie.Value)

	userID := services.GetUserIDBySession(cookie.Value)
	var req struct {
		AnimeID string `json:"animeID"`
		Status  string `json:"status"`
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Некорректный запрос", http.StatusBadRequest)
		return
	}

	log.Printf("Received anime_id: %q", req.AnimeID)

	cleanedID := strings.TrimSpace(req.AnimeID)
	animeID, err := strconv.Atoi(cleanedID)
	if err != nil {
		log.Printf("Failed to convert anime_id: %v", err)
		http.Error(w, "Некорректный anime_id", http.StatusBadRequest)
		return
	}

	log.Printf("Parsed anime_id: %d", animeID)
	err = services.UpdateAnimeStatus(userID, animeID, req.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("userID из сессии: %d", userID)

	log.Printf("UpdateAnimeStatus вызван с параметрами: userID=%d, animeID=%s, status=%q", userID, req.AnimeID, req.Status)

	err = services.UpdateAnimeStatus(userID, animeID, req.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func AnimeListHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := services.GetUserIDFromSession(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	log.Printf("Получен userID: %d", userID)
	log.Println("AnimeListHandler: вызов services.GetAnimesForUser")
	animes, err := services.GetAnimesForUser(userID)
	if err != nil {
		http.Error(w, "Ошибка GetAnimesForUser:", http.StatusInternalServerError)
		return
	}
	log.Printf("Получено аниме: %d", len(animes))
	log.Printf("Возвращаемые данные: %v", animes)

	err = templates.ExecuteTemplate(w, "index.html", map[string]interface{}{
		"Animes": animes,
	})
	if err != nil {
		log.Printf("Ошибка рендера шаблона: %v", err)
		http.Error(w, "Ошибка отображения", http.StatusInternalServerError)
	}
}
