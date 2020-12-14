package envrc

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFile(t *testing.T) {
	dir, err := ioutil.TempDir("", "test")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	tmpfn := filepath.Join(dir, ".envrc")

	f := NewFile(dir)
	p := f.Path
	assert.Equal(t, tmpfn, p, fmt.Sprintf("Want %s, got %s", tmpfn, p))
}
