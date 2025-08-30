package services

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

type FileService struct{}

func NewFileService() * FileService {
	return &FileService{}
}


// GenerateSHA256 takes a file path and returns the SHA-256 hash of its content.
// The hash is returned as a hex-encoded string. This is ideal for uniquely
// identifying a file based on its contents.
func (fs *FileService) GetFileHash(path string)(string, error){
	// Attempt to open the file from the given path.
	// os.Open returns the file and an error if the file cannot be found or accessed.
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	// defer ensures that the file is closed when the function exits,
	// regardless of whether it completes successfully or returns an error.
	defer file.Close()

	// Create a new SHA-256 hash object.
	// sha256.New() returns an object that implements the hash.Hash interface.
	hash := sha256.New()

	// Copy the file's content into the hash object.
	// io.Copy reads from the file and writes to the hash object.
	// This is memory-efficient as it reads the file in chunks, not all at once.
	if _, err := io.Copy(hash, file); err != nil {
		// If there's an error during the copy process, return an empty string and the error.
		return "", fmt.Errorf("failed to copy file content to hash: %w", err)
	}

	// Calculate the final hash sum.
	// hash.Sum(nil) computes the checksum and returns it as a byte slice.
	// We pass 'nil' because we want a new slice to be allocated for the result.
	hashInBytes := hash.Sum(nil)

	// Convert the byte slice to a hexadecimal string.
	// hex.EncodeToString makes the hash readable and easy to store.
	hashString := hex.EncodeToString(hashInBytes)

	log.Default().Printf("[GETFILEHASH] Generated hash: %s for file: %s", hashString, path)

	// Return the resulting hash string and no error.
	return hashString, nil
}