package main

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"journal-cli/src/models"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	dbDir := filepath.Join(homeDir, ".journal-cli")
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return err
	}
	dbPath := filepath.Join(dbDir, "journal.db")
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS entries (
		id TEXT PRIMARY KEY,
		added_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		text TEXT
	)`)
	return err
}

func generateUUID() (string, error) {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func addEntryToDB(entry string) (string, error) {
	id, err := generateUUID()
	if err != nil {
		return "", err
	}

	stmt, err := db.Prepare("INSERT INTO entries (id, text) VALUES (?, ?)")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, entry)
	if err != nil {
		return "", err
	}

	return id, nil
}

func getEntriesFromDB() ([]models.Entry, error) {
	rows, err := db.Query("SELECT id, added_at, updated_at, text FROM entries ORDER BY added_at ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []models.Entry
	for rows.Next() {
		var entry models.Entry
		err := rows.Scan(&entry.ID, &entry.AddedAt, &entry.UpdatedAt, &entry.Text)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func editEntryInDB(id string, newEntry string) (int64, error) {
	stmt, err := db.Prepare("UPDATE entries SET text = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(newEntry, id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func deleteEntryFromDB(id string) (int64, error) {
	stmt, err := db.Prepare("DELETE FROM entries WHERE id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
