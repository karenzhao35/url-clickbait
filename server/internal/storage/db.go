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

func queryData(conn *pgx.Conn) {
	rows, err := conn.Query(context.Background(), "SELECT * FROM urls")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var oldUrl string
		var newUrl string
		var createdAt time.Time
		var expiresAt time.Time
		err := rows.Scan(&id, &oldUrl, &newUrl, &createdAt, &expiresAt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Old URL: %s, New URL: %s, Created At: %s, Expires At: %s\n", id, oldUrl, newUrl, createdAt, expiresAt)
	}
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
	query := "SELECT old_url FROM urls WHERE new_url = $1"
	err = conn.QueryRow(context.Background(), query, key).Scan(&originalUrl)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", errors.New("url not found")
		}
		return "", fmt.Errorf("failed to query url: %w", err)
	}

	return originalUrl, nil
}

// TestDB tests the database connection and queries
func TestDB() error {
	conn, err := connect()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("Successfully connected to database!")

	// Check table structure
	if err := TestTableStructure(); err != nil {
		fmt.Printf("Error checking table structure: %v\n", err)
	}

	queryData(conn)
	return nil
}

// TestTableStructure tests what columns exist in the urls table
func TestTableStructure() error {
	conn, err := connect()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	// Query to see actual column names
	rows, err := conn.Query(context.Background(), "SELECT column_name FROM information_schema.columns WHERE table_name = 'urls'")
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Columns in 'urls' table:")
	for rows.Next() {
		var columnName string
		err := rows.Scan(&columnName)
		if err != nil {
			return err
		}
		fmt.Printf("- %s\n", columnName)
	}

	return nil
}
