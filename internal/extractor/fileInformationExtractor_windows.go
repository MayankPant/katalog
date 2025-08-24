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

// GetFileMetadata extracts metadata (Windows)
func GetFileMetadata(path string) (*FileMetadata, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	stat := info.Sys().(*syscall.Win32FileAttributeData)

	// Windows timestamps are in 100-nanosecond intervals since 1601
	nanoToTime := func(ft int64) time.Time {
		return time.Unix(0, ft*100)
	}

	ctime := nanoToTime(stat.CreationTime.Nanoseconds())
	atime := nanoToTime(stat.LastAccessTime.Nanoseconds())
	mtime := nanoToTime(stat.LastWriteTime.Nanoseconds())

	return &FileMetadata{
		Path:         path,
		Size:         info.Size(),
		IsDir:        info.IsDir(),
		Mode:         info.Mode(),
		LastModified: mtime,
		CreatedAt:    ctime,
		AccessedAt:   atime,
	}, nil
}