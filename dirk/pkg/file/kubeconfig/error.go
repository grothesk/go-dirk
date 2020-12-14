package kubeconfig

// ProcessKubeconfigError is thrown if processing kubeconfig fails
type ProcessKubeconfigError struct {
	Err error
}

func (e *ProcessKubeconfigError) Error() string {
	return "processing kubeconfig fails: " + e.Err.Error()
}

// CopyConfigfileError is thrown if copying configfile to kubeconfig fails
type CopyConfigfileError struct {
	Err error
}

func (e *CopyConfigfileError) Error() string {
	return "copying configfile to kubeconfig fails: " + e.Err.Error()
}

// SetModeError is thrown if setting mode of the kubeconfig fails
type SetModeError struct {
	Err error
}

func (e *SetModeError) Error() string {
	return "setting mode of kubeconfig file fails: " + e.Err.Error()
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
