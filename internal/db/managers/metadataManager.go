package managers

import (
	"database/sql"
	"fmt"
	"katalog/internal/extractor"
	"strconv"
)

type Metadata struct {
	ID int `json:"id"`
	FileID int `json:"file_id"`
	Key string `json:"key"`
	Value string `json:"value"`
}

type MetadataManager interface {
	SetupTable() error
	Create(metadata *Metadata) error
	InsertFileMetadata(file *File, fileMetadata *extractor.FileMetadata) error
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

func (m *metadataManager) Create(metadata *Metadata) error{
	query := fmt.Sprintf(`

	INSERT INTO %s (file_id, key, value) VALUES (?, ?, ?)
	
	`, m.tableName)

	_, err := m.DB.Exec(query, metadata.FileID, metadata.Key, metadata.Value)

	if err != nil {
		return fmt.Errorf("failed to create metadata entry: %w", err)
	}
	return nil
}

func (metadataManager *metadataManager) InsertFileMetadata(file *File, fileMetadata *extractor.FileMetadata) error {

	var metadataData []Metadata

	metadataData = append(metadataData, Metadata{
			FileID: file.ID,
			Key: "size",
			Value: strconv.Itoa(int(fileMetadata.Size)),
		})
	metadataData = append(metadataData, Metadata{
			FileID: file.ID,
			Key: "isDir",
			Value: strconv.FormatBool(fileMetadata.IsDir),
		})
	metadataData = append(metadataData, Metadata{
			FileID: file.ID,
			Key: "mode",
			Value: fileMetadata.Mode.String(),
		})
	metadataData = append(metadataData, Metadata{
			FileID: file.ID,
			Key: "lastModified",
			Value: fileMetadata.LastModified.String(),
		})
	metadataData = append(metadataData, Metadata{
			FileID: file.ID,
			Key: "createdAt",
			Value: fileMetadata.CreatedAt.String(),
		})
	metadataData = append(metadataData, Metadata{
			FileID: file.ID,
			Key: "accessedAt",
			Value: fileMetadata.AccessedAt.String(),
		})

	for _, value := range(metadataData) {
		err := metadataManager.Create(&value)
		if err != nil {
			fmt.Printf("Metadata insert error: %v\n", err)
        	return err
		}
	}
	return nil
	
}