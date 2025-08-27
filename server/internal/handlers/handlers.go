package handlers

import (
	"os"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/karenzhao35/shrimpy/internal/shortener"
	"github.com/karenzhao35/shrimpy/internal/storage"
)

type UrlHandler struct {
}

func (h UrlHandler) CreateUrl(w http.ResponseWriter, req *http.Request) {
	var requestData struct {
		OldUrl    string `json:"old_url"`
		ExpiredAt string `json:"expired_at"`
	}
	if err := json.NewDecoder(req.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	if requestData.OldUrl == "" {
		http.Error(w, "Old URL cannot be empty", http.StatusBadRequest)
		return
	}

	expiry, err := strconv.Atoi(requestData.ExpiredAt)

	if err != nil || expiry <= 0 {
		expiry = 30
	}

	key := shortener.GenerateCustomKey()

	err = storage.Save(key, requestData.OldUrl, expiry)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	response := map[string]string{"short_url": os.Getenv("BASE_URL") + "/make-money-fast/" + key}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h UrlHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")

	if key == "" {
		http.Error(w, "key is missing", http.StatusBadRequest)
		return
	}

	originalUrl, err := storage.GetOriginalUrl(key)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	http.Redirect(w, r, originalUrl, http.StatusMovedPermanently)
}