package envrc

// CountExportError is thrown if counting exports fails
type CountExportError struct {
	Err error
}

func (e *CountExportError) Error() string {
	return "cannot count exports: " + e.Err.Error()
}

// InvalidCountError is thrown if count of export is invalid
type InvalidCountError struct{}

func (e *InvalidCountError) Error() string {
	return "invalid export count"
}

// WriteExportError is thrown if count of export is invalid
type WriteExportError struct {
	Err error
}

func (e *WriteExportError) Error() string {
	return "cannot write export: " + e.Err.Error()
}
