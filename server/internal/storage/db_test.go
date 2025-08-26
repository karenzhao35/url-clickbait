package storage

import (
    "testing"
)

func TestDatabaseConnection(t *testing.T) {
    err := TestDB()
    if err != nil {
        t.Fatalf("Database test failed: %v", err)
    }
}
