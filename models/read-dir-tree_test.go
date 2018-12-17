package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"visual-file-server/utils"
)

func TestReadDirTree(t *testing.T) {
	return
	data, err := GetDirTree("../")
	if err != nil {
		t.Error(err)
	}
	fd, err := os.Create("log.json")
	defer fd.Close()
	d, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}
	fd.Write(d)
}

func TestFileCompress(t *testing.T) {
	file, _ := os.Open("../utils/compress.go")
	d, err := ioutil.ReadAll(file)
	t.Log(string(d))
	if err != nil {
		t.Error(err)
	}

	data, err := utils.FileCompress(file)

	if err != nil {
		t.Error(err)
	}

	t.Log(string(data))
}
