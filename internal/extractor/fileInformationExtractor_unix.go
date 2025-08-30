//go:build unix
// +build unix

package extractor

import (
	"os"
	"syscall"
	"time"
)

// FileMetadata holds basic metadata about a file
type FileMetadata struct {
	Path         string
	Size         int64
	IsDir        bool
	Mode         os.FileMode
	LastModified time.Time
	CreatedAt    time.Time
	AccessedAt   time.Time
}
// GetFileMetadata extracts metadata (Unix / Linux / macOS)
func GetFileMetadata(path string) (*FileMetadata, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	stat := info.Sys().(*syscall.Stat_t)

	// Access time
	atime := time.Unix(int64(stat.Atimespec.Sec), int64(stat.Atimespec.Nsec))
	// Birth time (not always available, sometimes ctime)
	ctime := time.Unix(int64(stat.Ctimespec.Sec), int64(stat.Ctimespec.Nsec))

	return &FileMetadata{
		Path:         path,
		Size:         info.Size(),
		IsDir:        info.IsDir(),
		Mode:         info.Mode(),
		LastModified: info.ModTime(),
		CreatedAt:    ctime,
		AccessedAt:   atime,
	}, nil
}
