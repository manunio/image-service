package files

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func setUpLocal(t *testing.T) (*Local, string, func()) {
	// create a temporary directory
	dir, err := ioutil.TempDir("", "files")
	if err != nil {
		t.Fatal(err)
	}

	l, err := NewLocal(dir, -1)
	if err != nil {
		t.Fatal(err)
	}
	return l, dir, func() {

	}
}

func TestSavesContentOfReader(t *testing.T) {
	savePath := "/1/test.png"
	fileContents := "Hello World!"
	l, dir, cleanup := setUpLocal(t)
	defer cleanup()

	err := l.Save(savePath, bytes.NewBuffer([]byte(fileContents)))
	assert.NoError(t, err)

	// check file has been correctly written
	f, err := os.Open(filepath.Join(dir, savePath))
	assert.NoError(t, err)

	// check the contents of file
	d, err := ioutil.ReadAll(f)
	assert.NoError(t, err)
	assert.Equal(t, fileContents, string(d))
}

func TestGetContentsAndWritesToWriter(t *testing.T) {
	savePath := "/1/test.png"
	fileContents := "Hello World"
	l, _, cleanup := setUpLocal(t)
	defer cleanup()

	// save a file
	err := l.Save(savePath, bytes.NewBuffer([]byte(fileContents)))
	assert.NoError(t, err)

	// Read the file back
	r, err := l.Get(savePath)
	assert.NoError(t, err)
	defer r.Close()

	// read the full contents of the reader
	d, err := ioutil.ReadAll(r)
	assert.Equal(t, fileContents, string(d))
}