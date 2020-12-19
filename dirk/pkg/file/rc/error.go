package rc

// SetupRcFileError is thrown if setting up rc file fails
type SetupRcFileError struct {
	Err error
}

func (e *SetupRcFileError) Error() string {
	return "setting up rc file fails: " + e.Err.Error()
}

// ReplaceOrAppendExportError is thrown if replacing or appending an export fails
type ReplaceOrAppendExportError struct {
	Err error
}

func (e *ReplaceOrAppendExportError) Error() string {
	return "replacing or appending export fails: " + e.Err.Error()
}

// CreateRcFileError is thrown if creating rc file fails
type CreateRcFileError struct {
	Err error
}

func (e *CreateRcFileError) Error() string {
	return "creating rc file fails: " + e.Err.Error()
}
