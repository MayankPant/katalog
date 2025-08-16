package managers

import (
	"database/sql"
	"fmt"
	"log"
)

type ConfigureDatabaseParams struct {
	DB		*sql.DB
}
type ConfigureDatabaseManager interface {
	ConfigureDatabase() error
}
type configureDatabaseManager struct {
	DB        *sql.DB
}

func NewConfigDatabaseManager(params ConfigureDatabaseParams) ConfigureDatabaseManager {
	return &configureDatabaseManager{
		DB: params.DB,
	}
}

func (m *configureDatabaseManager) ConfigureDatabase() error {
	query := `
	-- Helpful index for speeding up metadata lookups
	CREATE INDEX IF NOT EXISTS idx_metadata_file_id ON metadata(file_id);
	CREATE INDEX IF NOT EXISTS idx_file_tags_tag_id ON file_tags(tag_id);
	`
	log.Default().Print("[CONFIGURE DATABASE] Adding indexing to metatdata table for faster lookups")
	_, err := m.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to configure database: %w", err)
	}
	return nil
}
