package fileinfo

import (
	"os"
	"path/filepath"
)

// TotalFilesInDirectory - find the total number of files in the directory
func TotalFilesInDirectory(path string) (int64, error) {
	var totalFiles int64

	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			totalFiles++
		}

		return err
	})

	return totalFiles, err
}
