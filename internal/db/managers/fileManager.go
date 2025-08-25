package managers

import (
	"database/sql"
	"fmt"
	"time"
)

type File struct {
	ID        int    `json:"id"`
	Path      string `json:"path"`
	Hash      string `json:"hash"`
	Size      int    `json:"size"`
	CreatedAt time.Time `json:"created_at"`
	ScannedAt time.Time `json:"scanned_at"`
}

type FileManager interface {
	SetupTable() error
	Insert(files *File) error
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
func (m *fileManager) Insert(file *File) error {
    fmt.Printf("[FILEMANAGERINSERT] adding file to database.")
    query := fmt.Sprintf(`
        INSERT OR IGNORE INTO %s (
            path, hash, size, scanned_at
        ) VALUES (?, ?, ?, ?);
    `, m.tableName)

	fmt.Printf("\n[INSERT] insertion query: %s\n", query)

    _, err := m.DB.Exec(query, file.Path, file.Hash, file.Size, file.ScannedAt)
     if err != nil{
        fmt.Printf("Insert error: %v\n", err) // Add this line for more info
        return fmt.Errorf("\nfailed to insert file: %w\n", err)
    }
    return  nil
}