package services

import (
	"katalog/internal/db/managers"
	"katalog/internal/db/migrations"
	"katalog/internal/extractor"
	"fmt"
)
// used to persist files
func PersistCurrentFile(path string) bool {
	fmt.Printf("[PERSISTCURRENTFILE] Adding file %s to db.", path)
	db := migrations.GetDB()
	// extract file information and file hash
	fileInformation, err := extractor.GetFileMetadata(path)
	if err != nil {
		fmt.Printf("[PERSISTCURRENTFILE] Could not extract file information: %v\n", err)
		return  false
	}
	fileService := NewFileService()
	fileHash, err := fileService.GetFileHash(path)
	if err != nil{
		fmt.Printf("[PERSISTCURRENTFILE] Could not calculate file hash: %v\n", err)
		return false
	}

	// create entries into file table and metadata tables
	fileManager := managers.NewFileManager(managers.FileManagerParams{DB: db})
	response := fileManager.Insert(&managers.File{
		Path: path,
		Hash: fileHash,
		Size: int(fileInformation.Size),
		CreatedAt: fileInformation.CreatedAt,
		ScannedAt: fileInformation.AccessedAt,
	})

	if response != nil {
		fmt.Printf("[PERSISTCURRENTFILE] file insertion failed: %s.", err)
		return  false
	}

	return true
}