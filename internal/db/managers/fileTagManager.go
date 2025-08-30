package managers

import (
    "database/sql"
    "fmt"
)

type FileTag struct {
    FileID int `json:"file_id"`
    TagID  int `json:"tag_id"`
}

type FileTagManager interface {
    SetupTable() error
}

type fileTagManager struct {
    tableName string
    DB        *sql.DB
}

type FileTagManagerParams struct {
    DB *sql.DB
}

func NewFileTagManager(params FileTagManagerParams) FileTagManager {
    return &fileTagManager{
        tableName: "file_tags",
        DB:        params.DB,
    }
}

func (m *fileTagManager) SetupTable() error {
    query := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            file_id INTEGER NOT NULL,
            tag_id INTEGER NOT NULL,
            PRIMARY KEY (file_id, tag_id),
            FOREIGN KEY (file_id) REFERENCES files(id) ON DELETE CASCADE,
            FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
        );
    `, m.tableName)

    _, err := m.DB.Exec(query)
    if err != nil {
        return fmt.Errorf("failed to create table: %w", err)
    }
    return nil
}