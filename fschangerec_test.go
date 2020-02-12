package main

import (
	"github.com/mjms3/fschangerec/tempdir"
	"testing"
	"time"
)

func TestFsChangeRec(t *testing.T) {
	d := tempdir.NewTempDir("fschangerec_test_")
	messages := make(chan []string)
	defer d.Close()
	go monitor(d.Path, messages)
	time.Sleep(time.Second)
	newFile := "test_file"
	d.Write(newFile, "blah")
	changedFiles := <-messages
	if len(changedFiles) != 1 {
		t.Errorf("Expected 1 changed file, got: %s", changedFiles)
	}
	if changedFiles[0] != newFile {
		t.Errorf("Expected %s to be changed, got: %s", changedFiles, newFile)
	}
}
