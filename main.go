package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/containers/image/v5/docker"
	"github.com/containers/image/v5/docker/reference"
	"github.com/go-chi/chi/v5"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	r := chi.NewRouter()
	ctx := context.Background()

	r.Get("/tags/{imageRepo}", func(w http.ResponseWriter, r *http.Request) {
		imageRepo := chi.URLParam(r, "imageRepo")
		logger.Info(fmt.Sprintf("GET /tags/%s", imageRepo))
		imageRepo, err := url.PathUnescape(imageRepo)
		if err != nil {
			logger.Error("Could not unescape image repo", "imageRepo", imageRepo, "error", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		tags, err := getTags(ctx, imageRepo)
		if err != nil {
			logger.Error("Could not get tags", "repo", imageRepo, "error", err)
			// TODO: Actually check what the error is here instead of assuming
			http.Error(w, fmt.Sprintf("Repository %s not found", imageRepo), 404)
			return
		}
		data, err := json.Marshal(tags)
		if err != nil {
			logger.Error("Could not marshal tags", "repo", imageRepo, "tags", tags, "error", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.Write(data)
	})

	logger.Info("Starting server", "port", 8080)
	http.ListenAndServe(":8080", r)
}

func getTags(ctx context.Context, imageRepo string) ([]string, error) {
	ref, err := reference.ParseNormalizedNamed(imageRepo)
	if err != nil {
		return []string{}, fmt.Errorf("Could not parse image name: %s", err)
	}
	imgRef, err := docker.NewReference(reference.TagNameOnly(ref))
	if err != nil {
		return []string{}, fmt.Errorf("Could not create docker reference: %s", err)
	}
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	tags, err := docker.GetRepositoryTags(ctx, nil, imgRef)
	if err != nil {
		return []string{}, fmt.Errorf("Could not get tags: %s", err)
	}
	return tags, nil
}
