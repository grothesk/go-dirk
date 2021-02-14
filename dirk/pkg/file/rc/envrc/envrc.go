package envrc

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/grothesk/go-dirk/dirk/internal/logging"
	"github.com/grothesk/go-dirk/dirk/pkg/direnv"
	"github.com/grothesk/go-dirk/dirk/pkg/file"
	"github.com/spf13/viper"
)

var filename string = ".envrc"

// File handles .envrc files
type File struct {
	Path string
}

// NewFile returns a new .envrc file
func NewFile(dir string) File {
	return File{
		Path: path.Join(dir, filename),
	}
}

// Exists checks if .envrc exists
func (f *File) Exists() bool {
	return file.Exists(f.Path)
}

// ReplaceOrAppendExport replaces or appends an export
func (f *File) ReplaceOrAppendExport() error {
	count, err := f.countExport()
	if err != nil {
		return &CountExportError{
			Err: err,
		}
	}

	if count == 0 {
		logging.Logger.Info("append export in " + f.Path)
		return f.appendExport()
	} else if count == 1 {
		logging.Logger.Info("replace export in " + f.Path)
		return f.replaceExport()
	} else {
		return &InvalidCountError{}
	}
}

// Create creates an .envrc file
func (f *File) Create() error {
	logging.Logger.Info("create " + f.Path)
	fObj, err := os.Create(f.Path)
	if err != nil {
		return &file.CreateFileError{
			Err: err,
		}
	}
	defer fObj.Close()

	_, err = fObj.WriteString(viper.GetString("export") + "\n")
	if err != nil {
		return &WriteExportError{
			Err: err,
		}
	}

	return nil
}

// Allow direnv allows the .envrc file
func (f *File) Allow() error {
	logging.Logger.Info("direnv allow " + f.Path)
	return direnv.AllowPath(f.Path)
}

func (f *File) countExport() (int, error) {
	count := 0
	lines, err := file.ReadLines(f.Path)
	if err != nil {
		return count, &file.ReadLinesError{
			Err: err,
		}
	}

	for _, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "export KUBECONFIG=") {
			count++
		}
	}

	return count, nil
}

func (f *File) appendExport() error {
	fObj, err := os.OpenFile(f.Path, os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return &file.OpenFileError{
			Err: err,
		}
	}
	defer fObj.Close()

	_, err = fObj.WriteString(fmt.Sprintf("\n%s\n", viper.GetString("export")))
	if err != nil {
		return &WriteExportError{
			Err: err,
		}
	}

	return nil
}

func (f *File) replaceExport() error {
	lines, err := file.ReadLines(f.Path)
	if err != nil {
		return &file.ReadLinesError{
			Err: err,
		}
	}

	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "export KUBECONFIG=") {
			lines[i] = viper.GetString("export")
			break
		}
	}

	err = file.WriteLines(lines, f.Path)
	if err != nil {
		return &file.WriteLinesError{
			Err: err,
		}
	}

	return nil
}
