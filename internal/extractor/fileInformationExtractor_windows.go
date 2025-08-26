package extractor

import (
	"os"
	"syscall"
	"time"
)
// This constant represents the number of 100-nanosecond intervals
// between the Windows epoch (January 1, 1601) and the Unix epoch (January 1, 1970).
const windowsEpochOffset int64 = 116444736000000000

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

	ctime := nanoToTime(stat.CreationTime)
	atime := nanoToTime(stat.LastAccessTime)
	mtime := nanoToTime(stat.LastWriteTime)

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
// nanoToTime converts a Windows FILETIME struct into a Go time.Time object.
// It handles the critical conversion from the Windows epoch (1601-01-01)
// to the Unix epoch (1970-01-01) and adjusts the time unit from
// 100-nanosecond intervals to standard nanoseconds.
func nanoToTime(ft syscall.Filetime) (time.Time) {
	/*
		The FILETIME struct is a 64-bit value split into two 32-bit fields:
		HighDateTime and LowDateTime. We need to combine them to get a single
		int64 representing the total number of 100-nanosecond intervals.
		The high part is shifted left by 32 bits and then combined with the low part.
	*/
	nanos := (int64(ft.HighDateTime) << 32) + int64(ft.LowDateTime)

	/*
		Subtract the offset between the two epochs. The 'windowsEpochOffset'
		constant accounts for the number of 100-nanosecond intervals from
		Jan 1, 1601, to Jan 1, 1970. Subtracting this value effectively shifts
		our starting point to the Unix epoch.
	*/
	nanos -= windowsEpochOffset

	/*
		Finally, we use time.Unix to create a time.Time object.
		time.Unix(sec, nsec) expects nanoseconds, but our 'nanos' variable is
		in 100-nanosecond intervals. We multiply by 100 to convert to nanoseconds.
	*/
	return time.Unix(0, nanos*100)
}