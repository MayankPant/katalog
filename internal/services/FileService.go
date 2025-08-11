package services

import "katalog/internal/filesystem"

type FileService struct{}

func NewFileService() * FileService {
	return &FileService{}
}


func (fs *FileService) ScanDirectory(path string)([] string, error){
	return filesystem.ScanDirectory(path)
}