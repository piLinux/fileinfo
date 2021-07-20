package fileinfo_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/pilinux/fileinfo"
)

// TestTotalSizeDirectory - test `TotalSizeDirectory` function
func TestTotalSizeDirectory(t *testing.T) {
	// Create temp directories
	tempDir, err := ioutil.TempDir("./", "tmp")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tempDir2, err := ioutil.TempDir(tempDir, "tmp2")
	if err != nil {
		t.Fatalf("failed to create temp dir2: %v", err)
	}
	defer os.RemoveAll(tempDir2)

	// Create temp files
	tempFile1, err := ioutil.TempFile(tempDir, "file1")
	if err != nil {
		t.Fatalf("failed to create temp file1: %v", err)
	}
	if err := tempFile1.Truncate(1e4); err != nil {
		t.Fatalf("failed to fill temp file1 with data: %v", err)
	}
	defer os.Remove(tempDir + "/" + tempFile1.Name())

	tempFile2, err := ioutil.TempFile(tempDir, "file2")
	if err != nil {
		t.Fatalf("failed to create temp file2: %v", err)
	}
	if err := tempFile2.Truncate(1e5); err != nil {
		t.Fatalf("failed to fill temp file2  with data: %v", err)
	}
	defer os.Remove(tempDir + "/" + tempFile2.Name())

	tempFile3, err := ioutil.TempFile(tempDir2, "file3")
	if err != nil {
		t.Fatalf("failed to create temp file3: %v", err)
	}
	if err := tempFile3.Truncate(1e6); err != nil {
		t.Fatalf("failed to fill temp file3 with data: %v", err)
	}
	defer os.Remove(tempDir2 + "/" + tempFile3.Name())

	tempFile4, err := ioutil.TempFile(tempDir2, "file4")
	if err != nil {
		t.Fatalf("failed to create temp file4: %v", err)
	}
	if err := tempFile4.Truncate(1e7); err != nil {
		t.Fatalf("failed to fill temp file4  with data: %v", err)
	}
	defer os.Remove(tempDir2 + "/" + tempFile4.Name())

	// Test the function
	// Find the total size of the directory including all sub-directories
	totalSizeDirectory, err := fileinfo.TotalSizeDirectory(tempDir)
	if err != nil {
		t.Fatalf("error executing TotalSizeDirectory function: %v", err)
	}
	if totalSizeDirectory != (1e4 + 1e5 + 1e6 + 1e7) {
		t.Fatalf("want: %v, got: %v", (1e4 + 1e5 + 1e6 + 1e7), totalSizeDirectory)
	}

	// Find the size of a sub-directory
	totalSizeDirectory2, err := fileinfo.TotalSizeDirectory(tempDir2)
	if err != nil {
		t.Fatalf("error executing TotalSizeDirectory function: %v", err)
	}
	if totalSizeDirectory2 != (1e6 + 1e7) {
		t.Fatalf("want: %v, got: %v", (1e6 + 1e7), totalSizeDirectory)
	}
}
