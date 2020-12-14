package kubeconfig

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/grothesk/go-dirk/dirk/pkg/file"
	"github.com/spf13/viper"
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

// Process creates or configures the kubeconfig file
func (kf *File) Process() error {
	var err error
	configfile := viper.GetString("configfile")
	mode := viper.GetString("mode")

	if file.Exists(kf.Path) {
		fmt.Printf("dirk: %s does already exist.\n", kf.Path)
		if mode == "skip" {
			if configfile == "" {
				fmt.Printf("dirk: skip writing empty file to existing kubeconfig.\n")
			} else {
				fmt.Printf("dirk: skip writing %s to existing kubeconfig.\n", configfile)
			}
		} else if mode == "replace" {
			if configfile == "" {
				fmt.Printf("dirk: replace existing kubeconfig by empty file.\n")
				err = kf.replaceByEmptyfile()
				if err != nil {
					return &ProcessKubeconfigError{
						Err: &ReplaceByEmptyfileError{
							Err: err,
						},
					}
				}
			} else {
				fmt.Printf("dirk: replace existing kubeconfig by %s.\n", configfile)
				err = kf.replaceByConfigfile()
				if err != nil {
					return &ProcessKubeconfigError{
						Err: &ReplaceByConfigfileError{
							Err: err,
						},
					}
				}
			}
			err = kf.setMode()
			if err != nil {
				return &ProcessKubeconfigError{
					Err: &SetModeError{
						Err: err,
					},
				}
			}
		}
	} else {
		fmt.Printf("dirk: %s does not exist.\n", kf.Path)
		if configfile == "" {
			fmt.Printf("dirk: create empty %s.\n", kf.Path)
			err = kf.create()
			if err != nil {
				return &ProcessKubeconfigError{
					Err: err,
				}
			}
		} else {
			fmt.Printf("dirk: write %s to kubeconfig.\n", configfile)
			err = kf.replaceByConfigfile()
			if err != nil {
				return &ProcessKubeconfigError{
					Err: &ReplaceByConfigfileError{
						Err: err,
					},
				}
			}
		}
		err = kf.setMode()
		if err != nil {
			return &ProcessKubeconfigError{
				Err: &SetModeError{
					Err: err,
				},
			}
		}
	}
	return nil
}

func (kf *File) create() error {
	fObj, err := os.Create(kf.Path)
	if err != nil {
		return &file.CreateFileError{
			Err: err,
		}
	}
	defer fObj.Close()

	return nil
}

func (kf *File) replaceByConfigfile() error {
	c := viper.GetString("configfile")

	cObj, err := os.Open(c)
	if err != nil {
		return &file.OpenFileError{
			Err: err,
		}
	}
	defer cObj.Close()

	fObj, err := os.Create(kf.Path)
	if err != nil {
		return &file.CreateFileError{
			Err: err,
		}
	}
	defer fObj.Close()

	_, err = io.Copy(fObj, cObj)
	if err != nil {
		return &CopyConfigfileError{
			Err: err,
		}
	}

	return nil
}

func (kf *File) replaceByEmptyfile() error {
	return kf.create()
}

func (kf *File) setMode() error {
	if err := os.Chmod(kf.Path, 0600); err != nil {
		return &ChmodError{
			Err: err,
		}
	}
	return nil
}
