package fileinfo

import (
	"os"
	"path/filepath"
)

// TotalSizeDirectory - find the total size of the directory
func TotalSizeDirectory(path string) (int64, error) {
	var size int64

	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			size += info.Size()
		}

		return err
	})

	return size, err
}
