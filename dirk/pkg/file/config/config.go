package config

// File interfaces a config file
type File interface {
	Exists() bool
	Skip()
	ReplaceByEmptyfile() error
	ReplaceByConfigfile(string) error
	Create() error
	SetMode() error
}

// SetupFile sets up a config file
func SetupFile(f File, c string, m string) error {
	var err error

	if f.Exists() {
		if m == "skip" {
			f.Skip()
		} else if m == "replace" {
			if c == "" {
				err = f.ReplaceByEmptyfile()
				if err != nil {
					return &SetupError{
						Err: &ReplaceByEmptyfileError{
							Err: err,
						},
					}
				}
			} else {
				err = f.ReplaceByConfigfile(c)
				if err != nil {
					return &SetupError{
						Err: &ReplaceByConfigfileError{
							Err: err,
						},
					}
				}
			}
			err = f.SetMode()
			if err != nil {
				return &SetupError{
					Err: &SetModeError{
						Err: err,
					},
				}
			}
		}
	} else {
		if c == "" {
			err = f.Create()
			if err != nil {
				return &SetupError{
					Err: err,
				}
			}
		} else {
			err = f.ReplaceByConfigfile(c)
			if err != nil {
				return &SetupError{
					Err: &ReplaceByConfigfileError{
						Err: err,
					},
				}
			}
		}
		err = f.SetMode()
		if err != nil {
			return &SetupError{
				Err: &SetModeError{
					Err: err,
				},
			}
		}
	}

	return nil
}
