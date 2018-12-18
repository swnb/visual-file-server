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
	var zlibBuffer bytes.Buffer

	zlibWriter, err := zlib.NewWriterLevel(&zlibBuffer, zlib.BestCompression)
	if err != nil {
		return nil, err
	}
	defer zlibWriter.Close()

	_, err = bufio.NewReader(fd).WriteTo(zlibWriter)
	if err != nil {
		return nil, err
	}
	return zlibBuffer.Bytes(), nil
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
