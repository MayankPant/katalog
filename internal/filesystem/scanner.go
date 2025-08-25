package filesystem

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"katalog/internal/services"
)
type ScannerService struct{}

func NewScannerService() * ScannerService {
	return &ScannerService{}
}
const NUM_WORKERS = 10

func (s *ScannerService) ScanDirectory(root string) ([]string, error) {
	var files []string
	paths := make(chan string, 100)   // files to process
	results := make(chan string, 100) // processed files

	var wg sync.WaitGroup
	wg.Add(NUM_WORKERS)

	// Workers
	for i := 0; i < NUM_WORKERS; i++ {
		go func() {
			defer wg.Done()
			for path := range paths {
				// Here you could do processing on the file
				results <- path
			}
		}()
	}

	// Walk directory
	var walkErr error
	go func() {
		walkErr = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				if os.IsPermission(err) {
					fmt.Printf("Permission denied: %s\n", path)
					return nil
				}
				fmt.Printf("Error reading %s: %v\n", path, err)
				return nil
			}
			if !d.IsDir() {
				paths <- path
			}
			return nil
		})
		close(paths) // done sending to workers
	}()

	// Close results after workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results and insert file metadata to the database
	for file := range results {
		res := services.PersistCurrentFile(file)
		if !res{
			fmt.Println("Could not add file")
		}
		files = append(files, file)
	}

	return files, walkErr
}
