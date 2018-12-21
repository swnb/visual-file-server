package utils

import (
	"os"
	"testing"
)

func check(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func TestCompress(t *testing.T) {
	file, err := os.Open("./compress.go")
	check(t, err)
	defer file.Close()
	data, err := FileCompress(file)
	check(t, err)
	t.Log(string(data))
}
