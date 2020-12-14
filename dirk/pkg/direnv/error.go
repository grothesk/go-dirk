package direnv

// DirenvAllowError is thrown if executing direnv allow fails
type DirenvAllowError struct {
	Err error
}

func (e *DirenvAllowError) Error() string {
	return "direnv allow fails: " + e.Err.Error()
}
