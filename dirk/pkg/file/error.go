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
