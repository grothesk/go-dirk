package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupFile(t *testing.T) {
	var err error
	var f *MockFile

	// config file already exists
	f = new(MockFile)

	// setup expectations
	f.On("Exists").Return(true)
	f.On("Skip").Return()
	f.On("ReplaceByEmptyfile").Return(nil)
	f.On("ReplaceByConfigfile", "configfile").Return(nil)
	f.On("Create").Return(nil)
	f.On("SetMode").Return(nil)

	// calling SetupFile
	err = SetupFile(f, "configfile", "skip")
	assert.Nil(t, err)
	f.AssertNumberOfCalls(t, "Skip", 1)
	f.AssertNumberOfCalls(t, "Exists", 1)
	f.AssertNumberOfCalls(t, "ReplaceByEmptyfile", 0)
	f.AssertNumberOfCalls(t, "ReplaceByConfigfile", 0)
	f.AssertNumberOfCalls(t, "Create", 0)
	f.AssertNumberOfCalls(t, "SetMode", 0)

	err = SetupFile(f, "configfile", "replace")
	assert.Nil(t, err)
	f.AssertNumberOfCalls(t, "Exists", 2)
	f.AssertNumberOfCalls(t, "ReplaceByConfigfile", 1)
	f.AssertNumberOfCalls(t, "SetMode", 1)

	err = SetupFile(f, "", "replace")
	assert.Nil(t, err)
	f.AssertNumberOfCalls(t, "Exists", 3)
	f.AssertNumberOfCalls(t, "ReplaceByEmptyfile", 1)
	f.AssertNumberOfCalls(t, "SetMode", 2)

	err = SetupFile(f, "", "skip")
	assert.Nil(t, err)
	f.AssertNumberOfCalls(t, "Exists", 4)
	f.AssertNumberOfCalls(t, "Skip", 2)

	// config file does not exist
	f = new(MockFile)

	// setup expectations
	f.On("Exists").Return(false)
	f.On("Skip").Return()
	f.On("ReplaceByEmptyfile").Return(nil)
	f.On("ReplaceByConfigfile", "configfile").Return(nil)
	f.On("Create").Return(nil)
	f.On("SetMode").Return(nil)

	err = SetupFile(f, "configfile", "skip")
	assert.Nil(t, err)
	f.AssertNumberOfCalls(t, "Exists", 1)
	f.AssertNumberOfCalls(t, "ReplaceByConfigfile", 1)
	f.AssertNumberOfCalls(t, "SetMode", 1)
	f.AssertNumberOfCalls(t, "ReplaceByEmptyfile", 0)
	f.AssertNumberOfCalls(t, "Create", 0)
	f.AssertNumberOfCalls(t, "Skip", 0)

	err = SetupFile(f, "configfile", "replace")
	assert.Nil(t, err)
	f.AssertNumberOfCalls(t, "Exists", 2)
	f.AssertNumberOfCalls(t, "ReplaceByConfigfile", 2)
	f.AssertNumberOfCalls(t, "SetMode", 2)

	err = SetupFile(f, "", "replace")
	assert.Nil(t, err)
	f.AssertNumberOfCalls(t, "Exists", 3)
	f.AssertNumberOfCalls(t, "Create", 1)
	f.AssertNumberOfCalls(t, "SetMode", 3)

	err = SetupFile(f, "", "skip")
	assert.Nil(t, err)
	f.AssertNumberOfCalls(t, "Exists", 4)
	f.AssertNumberOfCalls(t, "Create", 2)
	f.AssertNumberOfCalls(t, "SetMode", 4)
}
