package migrations

import (
	"database/sql"
	"log"
	"katalog/internal/db/managers"

	_ "github.com/mattn/go-sqlite3"
)
// used to initialize database and configure its properties.
func IntializeDatabase(dbPath string) bool {
	db, err := sql.Open("sqlite3", dbPath);

	if(err != nil){
		log.Fatalf("Failed to open database: %v", err)
		return false
	}
	defer db.Close()

	// enable foreign keys in sqlite3. Important for cascading
	_, err = db.Exec(`PRAGMA foreign_keys=ON;`)
	if err != nil {
		log.Fatalf("Failed to enable foreign keys: %v", err)
		return false
	}

	// Create FileManagerParams with the db connection
    fileManagerParams := managers.FileManagerParams{
        DB: db,
    }
    // Instantiate Files table
    fileManager := managers.NewFileManager(fileManagerParams)
    if err := fileManager.SetupTable(); err != nil {
        log.Fatalf("Failed to setup files table: %v", err)
        return false
    }
	// instantiate Metadata table
	metadataManagerParams := managers.MetadataManagerParams{
		DB: db,
	}
	metadataManager := managers.NewMetadataManager(metadataManagerParams)
	if err := metadataManager.SetupTable(); err != nil {
		log.Fatalf("Failed to setup metadata table: %v", err)
        return false
	}

	// tag Manager
	tagManagerParams := managers.TagManagerParams{
		DB: db,
	}
	tagManager := managers.NewTagManager(tagManagerParams)
	if err := tagManager.SetupTable(); err != nil {
		log.Fatalf("Failed to setup tags table: %v", err)
        return false
	}


	// file tag Manager
	fileTagManagerParams := managers.FileTagManagerParams{
		DB: db,
	}
	fileTagManager := managers.NewFileTagManager(fileTagManagerParams)
	if err := fileTagManager.SetupTable(); err != nil {
		log.Fatalf("Failed to setup file_tag table: %v", err)
        return false
	}
	// finally  configure database manager
	configureDatabaseManagerParams := managers.ConfigureDatabaseParams{
		DB: db,
	}
	configureDatabaseManager := managers.NewConfigDatabaseManager(configureDatabaseManagerParams)
	if err := configureDatabaseManager.ConfigureDatabase(); err != nil {
		log.Fatalf("Failed to configure database: %v", err)
        return false
	}
    return true
}