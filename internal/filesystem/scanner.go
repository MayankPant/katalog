package filesystem

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// recursively scans through directories and shares the filePaths

func ScanDirectory(root string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil{
			if os.IsPermission(err){
				fmt.Printf("Error reading %s: %v\n", path, err)
				return nil // skip this directory
			}
			fmt.Printf("Error reading %s: %v\n", path, err)
            return nil // still skip
		}
		if !d.IsDir(){
			files = append(files, path)
		}
		return nil

	})

	return files, err
}