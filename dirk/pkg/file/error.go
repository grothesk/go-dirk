package file

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

// CopyFileError is thrown if copying a file fails
type CopyFileError struct {
	Err error
}

func (e *CopyFileError) Error() string {
	return "cannot copy file: " + e.Err.Error()
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

// SetupError is thrown if setting up configfile fails
type SetupError struct {
	Err error
}

func (e *SetupError) Error() string {
	return "setting up configfile fails: " + e.Err.Error()
}

// CopyConfigfileError is thrown if copying configfile to kubeconfig fails
type CopyConfigfileError struct {
	Err error
}

func (e *CopyConfigfileError) Error() string {
	return "copying configfile to kubeconfig fails: " + e.Err.Error()
}

// SetModeError is thrown if setting mode of configfile fails
type SetModeError struct {
	Err error
}

func (e *SetModeError) Error() string {
	return "setting mode of configfile file fails: " + e.Err.Error()
}

// ChmodError is thrown if chmod fails
type ChmodError struct {
	Err error
}

func (e *ChmodError) Error() string {
	return "chmod of a file fails: " + e.Err.Error()
}

// ReplaceByEmptyfileError is thrown if chmod fails
type ReplaceByEmptyfileError struct {
	Err error
}

func (e *ReplaceByEmptyfileError) Error() string {
	return "replacing kubeconfig by empty file fails: " + e.Err.Error()
}

// ReplaceByConfigfileError is thrown if chmod fails
type ReplaceByConfigfileError struct {
	Err error
}

func (e *ReplaceByConfigfileError) Error() string {
	return "replacing kubeconfig by configfile fails: " + e.Err.Error()
}
