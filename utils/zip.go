package utils

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"io"
	"os"
)

// Compress compress file by write buffer region of it
func Compress(fd *os.File) ([]byte, error) {
	var zlipBuffer bytes.Buffer
	zlibWriter, err := zlib.NewWriterLevel(&zlipBuffer, zlib.BestCompression)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(zlibWriter, bufio.NewReader(fd))
	if err != nil {
		return nil, err
	}
	return zlipBuffer.Bytes(), nil
}

func Compress1(fd *os.File) ([]byte, error) {
	var zlipBuffer bytes.Buffer
	var err error
	if zlibWriter, err = zlib.NewWriterLevel(&zlipBuffer, zlib.BestCompression); err == nil {
		if _, err = io.Copy(zlibWriter, bufio.NewReader(fd)); err == nil {
			return zlipBuffer.Bytes(), nil
		}
	}
	return nil, err
}
