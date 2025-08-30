package managers

import (
    "database/sql"
    "fmt"
)

type Tag struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

type TagManager interface {
    SetupTable() error
}

type tagManager struct {
    tableName string
    DB        *sql.DB
}

type TagManagerParams struct {
    DB *sql.DB
}

func NewTagManager(params TagManagerParams) TagManager {
    return &tagManager{
        tableName: "tags",
        DB:        params.DB,
    }
}

func (m *tagManager) SetupTable() error {
    query := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL UNIQUE
        );
    `, m.tableName)

    _, err := m.DB.Exec(query)
    if err != nil {
        return fmt.Errorf("failed to create table: %w", err)
    }
    return nil
}