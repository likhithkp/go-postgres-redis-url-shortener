package services

import (
	"context"
	"errors"
	"fmt"
	"go-postgres-redis-url-shortener/config"
	"log"

	"github.com/jackc/pgx"
)

type CreateCodeMessage struct {
	Message string
	Code    string
}

type FullUrlMessage struct {
	Url  string
	Code string
}

func InsertUrl(url string, code string) (message *CreateCodeMessage) {
	var urlCode string
	query := `INSERT INTO urls (short_url, long_url) VALUES ($1, $2) RETURNING short_url`
	err := config.DB.QueryRow(context.Background(), query, code, url).Scan(&urlCode)
	if err != nil {
		log.Printf("Error inserting URL: %v", err)
	}

	log.Println("URL successfully inserted!")
	msg := CreateCodeMessage{
		Message: "URL shortened",
		Code:    urlCode,
	}

	return &msg
}

func GetUrlByCode(code string) (*FullUrlMessage, error) {
	var url, urlCode string
	query := `SELECT long_url, short_url FROM urls WHERE short_url=$1`
	row := config.DB.QueryRow(context.Background(), query, code)
	err := row.Scan(&url, &urlCode)

	if errors.Is(err, pgx.ErrNoRows) {
		log.Println("No URL found for the given code:", code)
		return nil, fmt.Errorf("no URL found for the given code")
	}

	if err != nil {
		log.Println("Error while getting full URL:", err)
		return nil, err
	}

	log.Println("URL fetched successfully!")
	return &FullUrlMessage{
		Url:  url,
		Code: urlCode,
	}, nil
}
