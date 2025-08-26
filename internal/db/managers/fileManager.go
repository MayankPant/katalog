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
	CreatedAt time.Time `json:"created_at"`
}

type FileManager interface {
	SetupTable() error
	Insert(files *File) (*File, error)
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
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`, m.tableName)

	_, err := m.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	return nil
}
func (m *fileManager) Insert(file *File) (*File, error) {
    fmt.Printf("[FILEMANAGERINSERT] adding file to database.\n")
    query := fmt.Sprintf(`
        INSERT INTO %s (path, hash)
        VALUES (?, ?)
        ON CONFLICT(path) DO UPDATE SET
            hash=excluded.hash
    `, m.tableName)

    fmt.Printf("\n[INSERT] insertion query: %s\n", query)

    _, err := m.DB.Exec(query, file.Path, file.Hash)
    if err != nil {
        fmt.Printf("Insert error: %v\n", err)
        return nil, fmt.Errorf("\nfailed to insert file: %w\n", err)
    }

    // Query the file back by its unique path
    selectQuery := fmt.Sprintf(`
        SELECT id, path, hash, created_at
        FROM %s
        WHERE path = ?
    `, m.tableName)

    row := m.DB.QueryRow(selectQuery, file.Path)
    var insertedFile File
    err = row.Scan(&insertedFile.ID, &insertedFile.Path, &insertedFile.Hash, &insertedFile.CreatedAt)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch inserted file: %w", err)
    }
    return &insertedFile, nil
}