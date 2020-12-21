package tests

import (
	"awesomeProject1/models"
	"awesomeProject1/routes"
	"bytes"
	"encoding/json"
	"github.com/Kamva/mgm"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func handlePostImage(t *testing.T) {
	data := models.Image{
		Id:   "1",
		Data: []byte("image_data"),
	}

	encoded, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", "/image", bytes.NewReader(encoded))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.HandleImage)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	image, _ := models.GetImage("1")
	assert.Equal(t, []byte("image_data"), image.Data)
	assert.Equal(t, "1", image.Id)
}

func handleGetImage(t *testing.T) {
	_ = models.CreateImage([]byte("image_data"), "1")

	req, err := http.NewRequest("GET", "/image/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.HandleImage)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	got := models.Image{}
	_ = json.NewDecoder(rr.Body).Decode(&got)

	assert.Equal(t, "1", got.Id)
	assert.Equal(t, []byte("image_data"), got.Data)
}

func TestImageHandler(t *testing.T) {
	_ = os.Setenv("DB_NAME", "images")
	_ = os.Setenv("DB_URI", "mongodb://localhost:27017")

	err := mgm.SetDefaultConfig(nil, os.Getenv("DB_NAME"), options.Client().ApplyURI(os.Getenv("DB_URI")))
	if err != nil {
		t.Fatal(err)
	}

	_ = mgm.Coll(&models.Image{}).Drop(mgm.Ctx())
	t.Run("handle_post_image", handlePostImage)

	_ = mgm.Coll(&models.Image{}).Drop(mgm.Ctx())
	t.Run("handle_get_image", handleGetImage)
}
