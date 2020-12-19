package kubeconfig

import (
	"path"

	"github.com/grothesk/go-dirk/dirk/pkg/file"
)

var filename string = "kubeconfig"

// File handles a kubeconfig file
type File struct {
	Path string
}

// NewFile returns a new kubeconfig file
func NewFile(dir string) File {
	return File{
		Path: path.Join(dir, filename),
	}
}

// Exists checks if kubeconfig file exists
func (f *File) Exists() bool {
	return file.Exists(f.Path)
}

// Skip does almost nothing
func (f *File) Skip() {
	// Do nothing or log something
}

// ReplaceByEmptyfile replaces kubeconfig file by an empty file
func (f *File) ReplaceByEmptyfile() error {
	return f.Create()
}

// ReplaceByConfigfile writes a given config file to the file path of a kubeconfig file
func (f *File) ReplaceByConfigfile(c string) error {
	return file.Replace(f.Path, c)
}

// Create creates a kubeconfig file
func (f *File) Create() error {
	return file.Create(f.Path)
}

// SetMode sets mode of kubeconfig file to 600
func (f *File) SetMode() error {
	return file.SetMode(f.Path, 0600)
}
