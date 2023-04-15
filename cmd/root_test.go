package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetCurrentDirectory(t *testing.T) {
	dir := getCurrentDirectory()
	if dir == "" {
		t.Error("getCurrentDirectory() returned an empty string")
	}
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Errorf("getCurrentDirectory() returned a non-existent directory: %s", dir)
	}
}

func TestIsValid(t *testing.T) {
	// Create a test file
	testFile := filepath.Join(os.TempDir(), "test_file.txt")
	defer os.Remove(testFile)

	// Test the function with a file that does not exist
	if !isValid(testFile) {
		t.Error("isValid() returned false for a file that does not exist")
	}

	// Test the function with a file that already exists
	if !isValid(testFile) {
		t.Error("isValid() returned false for a file that already exists")
	}

	// Test the function with a directory
	testDir := os.TempDir()
	if isValid(testDir) {
		t.Error("isValid() returned true for a directory")
	}
}

func TestCheckIfExists(t *testing.T) {
	// Create a test file
	testFile := filepath.Join(os.TempDir(), "test_file.txt")
	defer os.Remove(testFile)

	// Test the function with a file that does not exist
	if checkIfExists(testFile) {
		t.Error("checkIfExists() returned true for a file that does not exist")
	}

	// Test the function with a file that already exists
	f, err := os.Create(testFile)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	if !checkIfExists(testFile) {
		t.Error("checkIfExists() returned false for a file that already exists")
	}

	// Test the function with a directory
	testDir := os.TempDir()
	if !checkIfExists(testDir) {
		t.Error("checkIfExists() returned false for an existing directory")
	}
}

func TestAppendToFileName(t *testing.T) {
	fp := "/path/to/file.txt"

	// Test the function with an empty string to append
	expected := fp
	got := appendToFileName(fp, "")
	if got != expected {
		t.Errorf("appendToFileName() with empty string did not return expected value. Expected: %s, Got: %s", expected, got)
	}

	// Test the function with a non-empty string to append
	expected = "/path/to/file (1).txt"
	got = appendToFileName(fp, " (1)")
	if got != expected {
		t.Errorf("appendToFileName() with non-empty string did not return expected value. Expected: %s, Got: %s", expected, got)
	}
}

func TestAppendTillNotExist(t *testing.T) {
	// Set up a test directory and file
	testDir := t.TempDir()
	testFile := filepath.Join(testDir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0666)

	// Test appending a string to the file
	expected := filepath.Join(testDir, "test (1).txt")
	result := appendTillNotExist(testFile)
	if result != expected {
		t.Errorf("appendTillNotExist(%s) = %s; expected %s", testFile, result, expected)
	}
	os.WriteFile(expected, []byte("test"), 0666)

	// Test appending a string to a file that already has an appended string
	testFile = filepath.Join(testDir, "test (1).txt")
	expected = filepath.Join(testDir, "test (1) (1).txt")
	result = appendTillNotExist(testFile)
	if result != expected {
		t.Errorf("appendTillNotExist(%s) = %s; expected %s", testFile, result, expected)
	}

	// Test appending a string to a file that already has multiple appended strings
	testFile = filepath.Join(testDir, "test.txt")
	expected = filepath.Join(testDir, "test (2).txt")
	result = appendTillNotExist(testFile)
	if result != expected {
		t.Errorf("appendTillNotExist(%s) = %s; expected %s", testFile, result, expected)
	}

	// Test appending a string to a file that doesn't exist yet
	testFile = filepath.Join(testDir, "new.txt")
	expected = testFile
	result = appendTillNotExist(testFile)
	if result != expected {
		t.Errorf("appendTillNotExist(%s) = %s; expected %s", testFile, result, expected)
	}
}
