package tempdir

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/scylladb/go-set/strset"
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
	fatalError(err)
	return d
}

func fatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (d *TemporaryDirectory) Write( path string, fileContent string) string {
	actualPath := filepath.Join(d.Path, path)

	f, err := os.Create(actualPath)
	defer f.Close()
	fatalError(err)
	_, err = f.WriteString(fileContent)
	fatalError(err)
	return actualPath
}

func (d *TemporaryDirectory) Compare(t *testing.T,expected []string) {
	files, err := ioutil.ReadDir(d.Path)
	fatalError(err)

	actualNames := strset.New()
	for _, f := range files {
		actualNames.Add(f.Name())
	}
	expectedNames := strset.New(expected...)
	inExpectedButNotActual := strset.Difference(expectedNames, actualNames)
	inActualButNotExpected := strset.Difference(actualNames, expectedNames)
	errorString := ""
	if !inExpectedButNotActual.IsEmpty(){
		errorString += fmt.Sprintf("In expected but not actual:\n%s\n", inExpectedButNotActual.String())
	}
	if !inActualButNotExpected.IsEmpty(){
		errorString += fmt.Sprintf("In actual but not expected:\n%s\n", inActualButNotExpected)
	}
	if len(errorString)>0{
		t.Errorf(errorString)
	}
}
func (d *TemporaryDirectory) Close( )  {
	err := os.RemoveAll(d.Path)
	fatalError(err)
}


