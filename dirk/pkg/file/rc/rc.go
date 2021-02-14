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
		err := f.ReplaceOrAppendExport()
		if err != nil {
			return &SetupRcFileError{
				Err: &ReplaceOrAppendExportError{
					Err: err,
				},
			}
		}
	} else {
		err := f.Create()
		if err != nil {
			return &SetupRcFileError{
				Err: &CreateRcFileError{
					Err: err,
				},
			}
		}
	}
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
