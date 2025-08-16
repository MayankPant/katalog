package managers

import (
	"database/sql"
	"fmt"
)

type Metadata struct {
	ID int `json:"id"`
	FileID int `json:"file_id"`
	Key string `json:"key"`
	Value string `json:"value"`
}

type MetadataManager interface {
	SetupTable() error
}

type metadataManager struct {
	tableName		string
	DB				*sql.DB		
}
type MetadataManagerParams struct {
	DB		*sql.DB
}

func NewMetadataManager(params MetadataManagerParams) MetadataManager{
	return &metadataManager{
		tableName: "metadata",
		DB: params.DB,
	}
}

func (m *metadataManager) SetupTable() error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		file_id INTEGER NOT NULL,
		key TEXT NOT NULL,
		value TEXT,
		FOREIGN KEY (file_id) REFERENCES files(id) ON DELETE CASCADE
	);`, m.tableName)

	_, err := m.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	return nil
}