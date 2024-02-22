package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	r := chi.NewRouter()

	r.Get("/tags/{imageRepo}", func(w http.ResponseWriter, r *http.Request) {
		imageRepo := chi.URLParam(r, "imageRepo")
		logger.Info(fmt.Sprintf("GET /tags/%s", imageRepo))
		tags := []string{"latest", "0.0.3", "0.0.2", "0.0.1"}
		data, err := json.Marshal(tags)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Write(data)
	})

	logger.Info("Starting server", "port", 8080)
	http.ListenAndServe(":8080", r)
}
