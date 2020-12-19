package config

// SetupError is thrown if setting up a config file fails
type SetupError struct {
	Err error
}

func (e *SetupError) Error() string {
	return "setting up config file fails: " + e.Err.Error()
}

// SetModeError is thrown if setting mode of a config file fails
type SetModeError struct {
	Err error
}

func (e *SetModeError) Error() string {
	return "setting mode of config file fails: " + e.Err.Error()
}

// ReplaceByEmptyfileError is thrown if chmod fails
type ReplaceByEmptyfileError struct {
	Err error
}

func (e *ReplaceByEmptyfileError) Error() string {
	return "replacing config file by empty file fails: " + e.Err.Error()
}

// ReplaceByConfigfileError is thrown if chmod fails
type ReplaceByConfigfileError struct {
	Err error
}

func (e *ReplaceByConfigfileError) Error() string {
	return "replacing config file by by another config file fails: " + e.Err.Error()
}
