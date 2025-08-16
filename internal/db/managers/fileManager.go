package managers

import (
	"time"
	"database/sql"
	"fmt"
)

type Files struct {
	ID        int    `json:"id"`
	Path      string `json:"path"`
	Hash      string `json:"string"`
	Size      int    `json:"size"`
	CreatedAt time.Time `json:"created_at"`
	ScannedAt time.Time `json:"scannned_at"`
}

type FileManager interface {
	SetupTable() error
}

type fileManager struct {
	tableName string
	DB        *sql.DB
}

type FileManagerParams struct {
	DB     *sql.DB
}

func NewFileManager(params FileManagerParams) FileManager{
	return &fileManager{
		tableName: "files",
		DB: params.DB,
	}
}

func (m *fileManager) SetupTable() error {
		query := fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			path TEXT NOT NULL UNIQUE,
			hash TEXT NOT NULL,
			size INTEGER NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			scanned_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`, m.tableName)

	_, err := m.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	return nil
}