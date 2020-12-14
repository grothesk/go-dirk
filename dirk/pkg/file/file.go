package file

import (
	"bufio"
	"fmt"
	"os"
)

// Exists checks if a file exists
func Exists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

// ReadLines reads lines from a text file and returns an array containing the lines
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, &OpenFileError{
			Err: err,
		}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

// WriteLines writes lines from a string array to a file
func WriteLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return &CreateFileError{
			Err: err,
		}
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
