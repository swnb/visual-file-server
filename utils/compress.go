package utils

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"io"
	"os"
)

// FileCompress compress file by write buffer region of it
func FileCompress(fd *os.File) ([]byte, error) {
	var zlipBuffer bytes.Buffer
	zlibWriter, err := zlib.NewWriterLevel(&zlipBuffer, zlib.BestCompression)
	if err != nil {
		return nil, err
	}
	defer zlibWriter.Close()

	_, err = io.Copy(zlibWriter, bufio.NewReader(fd))
	if err != nil {
		return nil, err
	}
	return zlipBuffer.Bytes(), nil
}

// DataCompress compress data by write buffer region of it
func DataCompress(data []byte) []byte {
	var zlipBuffer bytes.Buffer

	zlibWriter := zlib.NewWriter(&zlipBuffer)
	defer zlibWriter.Close()

	zlibWriter.Write(data)
	return zlipBuffer.Bytes()
}

// Decode deocde data with zlib
func DecodeData(data io.Reader) (io.ReadCloser, error) {
	return zlib.NewReader(data)
}
