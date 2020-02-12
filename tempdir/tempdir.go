package tempdir

import (
	"github.com/mjms3/fschangerec/comparisons"
	"github.com/mjms3/fschangerec/errorhandling"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

type TemporaryDirectory struct {
	Prefix string
	Path   string
}

const EMPTY_STRING string = ``

func NewTempDir(prefix string) *TemporaryDirectory {
	var err error
	d := new(TemporaryDirectory)
	d.Path, err = ioutil.TempDir(EMPTY_STRING, prefix)
	errorhandling.FatalError(err)
	return d
}

func (d *TemporaryDirectory) Write(path string, fileContent string) string {
	actualPath := filepath.Join(d.Path, path)

	f, err := os.Create(actualPath)
	defer f.Close()
	errorhandling.FatalError(err)
	_, err = f.WriteString(fileContent)
	errorhandling.FatalError(err)
	return actualPath
}

func (d *TemporaryDirectory) Compare(t *testing.T, expected []string) {
	files, err := ioutil.ReadDir(d.Path)
	errorhandling.FatalError(err)

	actualNames := make([]string, len(files))
	for idx, f := range files {
		actualNames[idx] = f.Name()
	}
	comparisons.CompareStringSlice(t, actualNames, expected)
}

func (d *TemporaryDirectory) Close() {
	err := os.RemoveAll(d.Path)
	errorhandling.FatalError(err)
}
