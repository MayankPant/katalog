package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

// recursively scans through directories and shares the filePaths

func ScanDirectory(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil{
			if os.IsPermission(err){
				fmt.Printf("Error reading %s: %v\n", path, err)
				return nil // skip this directory
			}
			fmt.Printf("Error reading %s: %v\n", path, err)
            return nil // still skip
		}
		if !info.IsDir(){
			files = append(files, path)
		}
		return nil

	})

	return files, err
}