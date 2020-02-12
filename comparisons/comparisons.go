package comparisons

import (
	"fmt"
	"github.com/scylladb/go-set/strset"
	"testing"
)

func CompareStringSlice(t *testing.T, actual []string, expected []string) {
	actualSet := strset.New(actual...)
	expectedSet := strset.New(expected...)
	inExpectedButNotActual := strset.Difference(expectedSet, actualSet)
	inActualButNotExpected := strset.Difference(actualSet, expectedSet)
	errorString := ""
	if !inExpectedButNotActual.IsEmpty() {
		errorString += fmt.Sprintf("In expected but not actual:\n%s\n", inExpectedButNotActual.String())
	}
	if !inActualButNotExpected.IsEmpty() {
		errorString += fmt.Sprintf("In actual but not expected:\n%s\n", inActualButNotExpected.String())
	}
	if len(errorString) > 0 {
		t.Errorf(errorString)
	}
}
