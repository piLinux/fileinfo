package fileinfo

import (
	"io/ioutil"
	"os"
	"testing"
)

// TestTotalFilesInDirectory - test `TotalFilesInDirectory` function
func TestTotalFilesInDirectory(t *testing.T) {
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
	defer os.Remove(tempDir + "/" + tempFile1.Name())

	tempFile2, err := ioutil.TempFile(tempDir, "file2")
	if err != nil {
		t.Fatalf("failed to create temp file2: %v", err)
	}
	defer os.Remove(tempDir + "/" + tempFile2.Name())

	tempFile3, err := ioutil.TempFile(tempDir2, "file3")
	if err != nil {
		t.Fatalf("failed to create temp file3: %v", err)
	}
	defer os.Remove(tempDir2 + "/" + tempFile3.Name())

	tempFile4, err := ioutil.TempFile(tempDir2, "file4")
	if err != nil {
		t.Fatalf("failed to create temp file4: %v", err)
	}
	defer os.Remove(tempDir2 + "/" + tempFile4.Name())

	// Test the function
	// Find the number of files in the directory including all sub-directories
	totalFilesInDirectory, err := TotalFilesInDirectory(tempDir)
	if err != nil {
		t.Fatalf("error executing TotalFilesInDirectory function: %v", err)
	}
	if totalFilesInDirectory != 4 {
		t.Fatalf("want: %v, got: %v", 4, totalFilesInDirectory)
	}

	// Find the number of files in a sub-directory
	totalFilesInDirectory2, err := TotalFilesInDirectory(tempDir2)
	if err != nil {
		t.Fatalf("error executing TotalFilesInDirectory function: %v", err)
	}
	if totalFilesInDirectory2 != 2 {
		t.Fatalf("want: %v, got: %v", 2, totalFilesInDirectory2)
	}
}
