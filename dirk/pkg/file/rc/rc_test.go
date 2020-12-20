package rc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupFile(t *testing.T) {
	var err error
	var f *MockFile

	// rc file already exists
	f = new(MockFile)

	// setup expectations
	f.On("Exists").Return(true)
	f.On("ReplaceOrAppendExport").Return(nil)
	f.On("Create").Return(nil)
	f.On("Allow").Return(nil)

	err = SetupFile(f)
	assert.Nil(t, err)
	f.AssertNumberOfCalls(t, "Exists", 1)
	f.AssertNumberOfCalls(t, "ReplaceOrAppendExport", 1)
	f.AssertNumberOfCalls(t, "Create", 0)
	f.AssertNumberOfCalls(t, "Allow", 1)

	// rc file does not exist
	f = new(MockFile)

	// setup expectations
	f.On("Exists").Return(false)
	f.On("ReplaceOrAppendExport").Return(nil)
	f.On("Create").Return(nil)
	f.On("Allow").Return(nil)

	err = SetupFile(f)
	assert.Nil(t, err)
	f.AssertNumberOfCalls(t, "Exists", 1)
	f.AssertNumberOfCalls(t, "ReplaceOrAppendExport", 0)
	f.AssertNumberOfCalls(t, "Create", 1)
	f.AssertNumberOfCalls(t, "Allow", 1)
}
