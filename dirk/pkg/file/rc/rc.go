package rc

import (
	"github.com/grothesk/go-dirk/dirk/pkg/direnv"
)

// File interfaces a rc file
type File interface {
	Exists() bool
	ReplaceOrAppendExport() error
	Create() error
	Allow() error
}

// SetupFile creates or configures a setup rc file
func SetupFile(f File) error {
	if f.Exists() {
		//fmt.Printf("dirk: %s does already exist.\n", ef.Path)
		//fmt.Printf("dirk: process %s.\n", ef.Path)
		err := f.ReplaceOrAppendExport()
		if err != nil {
			return &SetupRcFileError{
				Err: &ReplaceOrAppendExportError{
					Err: err,
				},
			}
		}
	} else {
		//fmt.Printf("dirk: %s does not exist.\n", ef.Path)
		//fmt.Printf("dirk: create %s.\n", ef.Path)
		err := f.Create()
		if err != nil {
			return &SetupRcFileError{
				Err: &CreateRcFileError{
					Err: err,
				},
			}
		}
	}
	//fmt.Printf("dirk: direnv allow %s.\n", ef.Path)
	err := f.Allow()
	if err != nil {
		return &SetupRcFileError{
			Err: &direnv.DirenvAllowError{
				Err: err,
			},
		}
	}

	return nil
}
