package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

func connect() (*pgx.Conn, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, errors.New("environmental variables not properly set")
	}

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func Save(key string, oldUrl string, expiry int) error {
	conn, err := connect()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	createdAt := time.Now()
	expiredAt := createdAt.Add(time.Duration(expiry) * time.Hour * 24)

	_, err = conn.Exec(context.Background(), `INSERT INTO urls ("old_url", "new_url", "created_at", "expires_at") VALUES ($1, $2, $3, $4)`,
		oldUrl, key, createdAt, expiredAt)
	if err != nil {
		log.Fatalf("Failed to save URL: %v", err)
	}

	return nil
}

func GetOriginalUrl(key string) (string, error) {
	conn, err := connect()
	if err != nil {
		return "", err
	}
	defer conn.Close(context.Background())

	var originalUrl string
	var expiresAt time.Time
	query := "SELECT old_url, expires_at FROM urls WHERE new_url = $1"
	err = conn.QueryRow(context.Background(), query, key).Scan(&originalUrl, &expiresAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", errors.New("url not found")
		}
		return "", fmt.Errorf("failed to query url: %w", err)
	}

	if time.Now().After(expiresAt) {
		_, deleteErr := conn.Exec(context.Background(), "DELETE FROM urls WHERE new_url = $1", key)
		if deleteErr != nil {
			log.Printf("Failed to delete expired URL %s: %v", key, deleteErr)
		}
		return "", errors.New("url has expired")
	}

	return originalUrl, nil
}
