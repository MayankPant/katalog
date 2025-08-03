package filesystem

import (
	"os"
	"path/filepath"
)

// recursively scans through directories and shares the filePaths

func ScanDirectory(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil{
			return err
		}
		if !info.IsDir(){
			files = append(files, path)
		}
		return nil

	})

	return files, err
}