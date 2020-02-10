package tempdir

import (
	"io/ioutil"
	"testing"
)

func TestWrite(t *testing.T) {
	d := NewTempDir("test_")
	defer d.Close()
	actualPath:= d.Write("first_test", "test file content")
	content, _ := ioutil.ReadFile(actualPath)
	contentString := string(content)
	expectedContent := "test file content"
	if contentString != expectedContent {
		t.Errorf("Read unexpected content from testfile: %s. Expected: %s, Actual: %s\n", actualPath, expectedContent, contentString)
	}
}

func TestCompareNothingPresent(t *testing.T){
	d := NewTempDir("test_")
	defer d.Close()
	expected := []string{}
	d.Compare(t,expected)
}

func TestCompareFilePresent(t *testing.T){
	d := NewTempDir("test_")
	defer d.Close()
	d.Write("file1","")
	expected := []string{"file1"}
	d.Compare(t,expected)
}
