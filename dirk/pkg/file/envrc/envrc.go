package envrc

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/grothesk/go-dirk/dirk/pkg/direnv"
	"github.com/grothesk/go-dirk/dirk/pkg/file"
	"github.com/spf13/viper"
)

var filename string = ".envrc"

// CountExportError is thrown if counting exports fails
type CountExportError struct {
	Err error
}

func (e *CountExportError) Error() string {
	return "cannot count exports: " + e.Err.Error()
}

// InvalidCountError is thrown if count of export is invalid
type InvalidCountError struct {
	Err error
}

func (e *InvalidCountError) Error() string {
	return "invalid export count: " + e.Err.Error()
}

// WriteExportError is thrown if count of export is invalid
type WriteExportError struct {
	Err error
}

func (e *WriteExportError) Error() string {
	return "cannot write export: " + e.Err.Error()
}

// ProcessEnvrcError is thrown if processing envrc fails
type ProcessEnvrcError struct {
	Err error
}

func (e *ProcessEnvrcError) Error() string {
	return "processing envrc fails: " + e.Err.Error()
}

// CreateEnvrcError is thrown if creating envrc fails
type CreateEnvrcError struct {
	Err error
}

func (e *CreateEnvrcError) Error() string {
	return "creating envrc fails: " + e.Err.Error()
}

// ReplaceOrAppendExportError is thrown if creating envrc fails
type ReplaceOrAppendExportError struct {
	Err error
}

func (e *ReplaceOrAppendExportError) Error() string {
	return "replacing or appendig export to envrc fails: " + e.Err.Error()
}

// File handles an .envrc file
type File struct {
	Path string
}

// NewFile returns a new .envrc file
func NewFile(dir string) File {
	return File{
		Path: path.Join(dir, filename),
	}
}

// Process creates or configures the .envrc file
func (ef *File) Process() error {
	if file.Exists(ef.Path) {
		fmt.Printf("dirk: %s does already exist.\n", ef.Path)
		fmt.Printf("dirk: process %s.\n", ef.Path)
		err := ef.replaceOrAppendExport()
		if err != nil {
			return &ProcessEnvrcError{
				Err: &ReplaceOrAppendExportError{
					Err: err,
				},
			}
		}
	} else {
		fmt.Printf("dirk: %s does not exist.\n", ef.Path)
		fmt.Printf("dirk: create %s.\n", ef.Path)
		err := ef.create()
		if err != nil {
			return &ProcessEnvrcError{
				Err: &CreateEnvrcError{
					Err: err,
				},
			}
		}
	}
	fmt.Printf("dirk: direnv allow %s.\n", ef.Path)
	ef.allow()

	return nil
}

func (ef *File) replaceOrAppendExport() error {
	count, err := ef.countExport()
	if err != nil {
		return &CountExportError{
			Err: err,
		}
	}

	if count == 0 {
		return ef.appendExport()
	} else if count == 1 {
		return ef.replaceExport()
	} else {
		return &InvalidCountError{
			Err: err,
		}
	}

	return nil
}

func (ef *File) countExport() (int, error) {
	count := 0
	lines, err := file.ReadLines(ef.Path)
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

func (ef *File) appendExport() error {
	fObj, err := os.OpenFile(ef.Path, os.O_WRONLY|os.O_APPEND, 0600)
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

func (ef *File) replaceExport() error {
	lines, err := file.ReadLines(ef.Path)
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

	err = file.WriteLines(lines, ef.Path)
	if err != nil {
		return &file.WriteLinesError{
			Err: err,
		}
	}

	return nil
}

func (ef *File) create() error {
	fObj, err := os.Create(ef.Path)
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

func (ef *File) allow() {
	direnv.AllowPath(ef.Path)
}
