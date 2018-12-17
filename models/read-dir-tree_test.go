package models

import (
	"encoding/json"
	"os"
	"testing"
)

func TestReadDirTree(t *testing.T) {
	data, err := GetDirTree("../")
	if err != nil {
		t.Error(err)
	}
	fd, err := os.Create("log")
	defer fd.Close()
	d, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}
	fd.Write(d)
}
