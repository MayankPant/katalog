package migrations

import (
	"database/sql"
	"log"
	"sync"
	"katalog/internal/db/managers"

	_ "github.com/mattn/go-sqlite3"
)

var (
	dbOnce sync.Once
	db *sql.DB
)
// used to initialize database and configure its properties.
// returns a pointer to the database for injection into various services
// used to be called once
func IntializeDatabase(dbPath string) (*sql.DB, error) {

	var err error
	dbOnce.Do(func() {
	db, err = sql.Open("sqlite3", dbPath);

	if(err != nil){
		log.Fatalf("Failed to open database: %v", err)
		return
	}

	// enable foreign keys in sqlite3. Important for cascading
	_, err = db.Exec(`PRAGMA foreign_keys=ON;`)
	if err != nil {
		log.Fatalf("Failed to enable foreign keys: %v", err)
		return
	}

	// Create FileManagerParams with the db connection
    fileManagerParams := managers.FileManagerParams{
        DB: db,
    }
    // Instantiate Files table
    fileManager := managers.NewFileManager(fileManagerParams)
    if err = fileManager.SetupTable(); err != nil {
        log.Fatalf("Failed to setup files table: %v", err)
		return
    }
	// instantiate Metadata table
	metadataManagerParams := managers.MetadataManagerParams{
		DB: db,
	}
	metadataManager := managers.NewMetadataManager(metadataManagerParams)
	if err = metadataManager.SetupTable(); err != nil {
		log.Fatalf("Failed to setup metadata table: %v", err)
		return
	}

	// tag Manager
	tagManagerParams := managers.TagManagerParams{
		DB: db,
	}
	tagManager := managers.NewTagManager(tagManagerParams)
	if err = tagManager.SetupTable(); err != nil {
		log.Fatalf("Failed to setup tags table: %v", err)
		return
	}


	// file tag Manager
	fileTagManagerParams := managers.FileTagManagerParams{
		DB: db,
	}
	fileTagManager := managers.NewFileTagManager(fileTagManagerParams)
	if err = fileTagManager.SetupTable(); err != nil {
		log.Fatalf("Failed to setup file_tag table: %v", err)
		return
	}
	// finally  configure database manager
	configureDatabaseManagerParams := managers.ConfigureDatabaseParams{
		DB: db,
	}
	configureDatabaseManager := managers.NewConfigDatabaseManager(configureDatabaseManagerParams)
	if err = configureDatabaseManager.ConfigureDatabase(); err != nil {
		log.Fatalf("Failed to configure database: %v", err)
		return
	}
})
	
	if err != nil {
		// If initialization failed, reset for a potential retry.
		// Be cautious with this in a real app.
		dbOnce = sync.Once{} 
		db = nil
		return nil, err
	}

	return db, nil
}


// GetDB is a convenience function to get the initialized pool.
// It will panic if InitDB has not been called successfully.
func GetDB() *sql.DB {
	if db == nil {
		panic("database has not been initialized")
	}
	return db
}