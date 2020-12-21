package routes

import (
	"awesomeProject1/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func handlePostImage(w http.ResponseWriter, r *http.Request) {
	var image models.Image

	err := json.NewDecoder(r.Body).Decode(&image)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = models.CreateImage(image.Data, image.Id)
	if err != nil {
		http.Error(w, "Failed to create image", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func handleGetImage(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/image/")

	image, err := models.GetImage(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get image with id %s", id), http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(bytes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleImage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		fmt.Println("POST /image")
		handlePostImage(w, r)

	case http.MethodGet:
		fmt.Println("GET /image")
		handleGetImage(w, r)

	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}
