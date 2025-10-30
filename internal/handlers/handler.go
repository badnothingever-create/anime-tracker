package handlers

import (
	"encoding/json"
	//"log"
	"net/http"

	"anime-tracker/internal/models"
	"anime-tracker/internal/repositories"
	"anime-tracker/internal/services"
)

// InitRoutes регистрирует роуты на переданном mux (явная маршрутизация)
func InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", LoginHandler)            // Страница входа
	mux.HandleFunc("/register", RegisterHandler) // Страница регистрации
	mux.HandleFunc("/logout", LogoutHandler)     // Выход
	mux.HandleFunc("/anime/status", UpdateAnimeStatusHandler)
	mux.HandleFunc("/anime", func(w http.ResponseWriter, r *http.Request) {
		userID, err := services.GetUserIDFromSession(r)
		if err != nil {
			// Например, редирект на страницу входа, если сессия невалидна
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		//log.Printf("DEBUG: userID из сессии: %d, для запроса %s", userID, r.URL.Path)
		switch r.Method {
		case http.MethodGet:

			animes, err := services.GetAnimesForUser(userID)
			if err != nil {
				http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
				return
			}
			//log.Printf("DEBUG: получено %d аниме для userID %d", len(animes), userID)
			username, err := repositories.GetUsernameByUserID(userID)
			err = templates.ExecuteTemplate(w, "index.html", map[string]interface{}{
				"Animes":   animes,
				"Username": username,
			})

			if err != nil {
				http.Error(w, "Ошибка шаблона: "+err.Error(), http.StatusInternalServerError)
				return
			}

		case http.MethodPost:
			var newAnime models.Anime
			err := json.NewDecoder(r.Body).Decode(&newAnime)
			if err != nil {
				http.Error(w, `{"error":"Некорректный запрос"}`, http.StatusBadRequest)
				return
			}
			err = services.CreateAnime(newAnime)
			if err != nil {
				http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)

		default:
			w.Header().Set("Allow", "GET, POST")
			http.Error(w, `{"error":"Метод не поддерживается"}`, http.StatusMethodNotAllowed)
		}
	})
}
