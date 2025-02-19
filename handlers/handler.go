package handlers

import (
	"encoding/json"
	"go-postgres-redis-url-shortener/helper"
	"go-postgres-redis-url-shortener/services"
	"net/http"
)

type RequestBody struct {
	Url string `json:"url"`
}

type CustomResponse struct {
	Message string
}

func GetFullUrl(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	message, err := services.GetUrlByCode(code)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}

func CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	var reqBody RequestBody

	if r.Body == nil {
		errorMessage := CustomResponse{
			Message: "URL cannot be empty",
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(&errorMessage)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil || reqBody.Url == "" {
		errorMessage := CustomResponse{
			Message: "Please enter the valid url",
		}
		json.NewEncoder(w).Encode(&errorMessage)
		return
	}

	code := helper.GenerateCode()
	message := services.InsertUrl(reqBody.Url, code)

	if message != nil {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(&message)
	}

}
