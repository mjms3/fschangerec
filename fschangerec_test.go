package main

import (
	"github.com/mjms3/fschangerec/comparisons"
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
	expected := []string{newFile}

	comparisons.CompareStringSlice(t, changedFiles, expected)
}
