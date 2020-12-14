package envrc

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
