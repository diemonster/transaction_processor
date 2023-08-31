package filereader

import (
	"os"
	"testing"
)

func TestReadLines(t *testing.T) {
	// Create a temporary file for testing
	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(file.Name())

	// Write some lines to the file
	lines := []string{"Line 1", "Line 2", "Line 3"}
	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			t.Fatalf("Failed to write line to file: %v", err)
		}
	}

	// Create a FileReader instance
	fr := NewFileReader(file.Name())

	// Read lines from the file
	readLines, err := fr.ReadLines()
	if err != nil {
		t.Fatalf("Failed to read lines from file: %v", err)
	}

	// Verify that the lines read match the expected lines
	i := 0
	for line := range readLines {
		if line != lines[i] {
			t.Errorf("Expected line '%s', got '%s'", lines[i], line)
		}
		i++
	}
}
