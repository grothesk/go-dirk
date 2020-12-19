package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Exists checks if a file exists
func Exists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

// Create creates a file
func Create(p string) error {
	f, err := os.Create(p)
	if err != nil {
		return &CreateFileError{
			Err: err,
		}
	}
	defer f.Close()

	return nil
}

// Replace replaces a file by another file
func Replace(dst string, src string) error {
	fs, err := os.Open(src)
	if err != nil {
		return &OpenFileError{
			Err: err,
		}
	}
	defer fs.Close()

	fd, err := os.Create(dst)
	if err != nil {
		return &CreateFileError{
			Err: err,
		}
	}
	defer fd.Close()

	_, err = io.Copy(fd, fs)
	if err != nil {
		return &CopyFileError{
			Err: err,
		}
	}

	return nil
}

// SetMode sets mode of a file
func SetMode(p string, m os.FileMode) error {
	if err := os.Chmod(p, m); err != nil {
		return &ChmodError{
			Err: err,
		}
	}
	return nil
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
