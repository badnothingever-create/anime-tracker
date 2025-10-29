package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"anime-tracker/internal/models"
	"anime-tracker/internal/services"
)

// InitRoutes регистрирует роуты на переданном mux (явная маршрутизация)
func InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", LoginHandler)                         // Страница входа
	mux.HandleFunc("/register", RegisterHandler)              // Страница регистрации
	mux.HandleFunc("/logout", LogoutHandler)                  // Выход
	mux.HandleFunc("/anime/status", UpdateAnimeStatusHandler) // Обновление статуса
	//mux.HandleFunc("/anime", AnimeListHandler)                // Страница с аниме
	mux.HandleFunc("/anime", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:

			animes, err := services.ListAnimes()
			if err != nil {
				http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
				return
			}

			for _, anime := range animes {
				log.Printf("Anime ID: %d, Title: %s, StatusString: %q", anime.ID, anime.Title, anime.StatusString())
			}

			err = templates.ExecuteTemplate(w, "index.html", map[string]any{
				"Animes": animes,
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
