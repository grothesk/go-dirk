package file

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExists(t *testing.T) {
	var exists bool
	var want bool

	dir, err := ioutil.TempDir("", "test")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	tmpfn := filepath.Join(dir, "test")

	exists = Exists(tmpfn)
	want = false
	assert.Equal(t, exists, want, fmt.Sprintf("Want %t, got %t", want, exists))

	file, err := os.Create(tmpfn)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	exists = Exists(tmpfn)
	want = true
	assert.Equal(t, exists, want, fmt.Sprintf("Want %t, got %t", want, exists))
}

func TestReadLines(t *testing.T) {
	var lines []string
	content := []byte("this\nis\na\ntest\n")
	contentArr := []string{"this", "is", "a", "test"}

	dir, err := ioutil.TempDir("", "test")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	_, err = ReadLines(filepath.Join(dir, "doesnotexist"))
	if assert.NotNil(t, err) {
		e, ok := err.(*OpenFileError)
		assert.Equal(t, ok, true, fmt.Sprintf("Want *file.CantOpenFileError, got %s", reflect.TypeOf(e)))
	}

	tmpfn := filepath.Join(dir, "test")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}

	lines, err = ReadLines(filepath.Join(dir, "test"))
	if err != nil {
		log.Fatal(err)
	}
	for i := range lines {
		assert.Equal(t, lines[i], contentArr[i], fmt.Sprintf("Want %s, got %s", contentArr[i], lines[i]))
	}
}

func TestWriteLines(t *testing.T) {
	wlines := []string{"this", "is", "a", "test"}

	dir, err := ioutil.TempDir("", "test")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	tmpfn := filepath.Join(dir, "test")
	err = WriteLines(wlines, tmpfn)

	rlines, err := ReadLines(tmpfn)
	if err != nil {
		log.Fatal(err)
	}
	for i := range wlines {
		if wlines[i] != rlines[i] {
			t.Fatalf("Want %s, got %s", wlines[i], rlines[i])
		}
	}
}
