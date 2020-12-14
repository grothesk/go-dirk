package file

import (
	"bufio"
	"fmt"
	"os"
)

// OpenFileError is thrown if a file cannot be open
type OpenFileError struct {
	Err error
}

func (e *OpenFileError) Error() string {
	return "cannot open file: " + e.Err.Error()
}

// CreateFileError is thrown if a file cannot be created
type CreateFileError struct {
	Err error
}

func (e *CreateFileError) Error() string {
	return "cannot create file: " + e.Err.Error()
}

// ReadLinesError is thrown if reading lines from a file fails
type ReadLinesError struct {
	Err error
}

func (e *ReadLinesError) Error() string {
	return "cannot read lines: " + e.Err.Error()
}

// WriteLinesError is thrown if writing lines to file fails
type WriteLinesError struct {
	Err error
}

func (e *WriteLinesError) Error() string {
	return "cannot write lines: " + e.Err.Error()
}

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
