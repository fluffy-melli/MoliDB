package gzip

import (
	"bytes"
	"compress/gzip"
	"io"
)

func Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)
	_, err := gzipWriter.Write(data)
	if err != nil {
		return nil, err
	}
	err = gzipWriter.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decompress(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(data)
	gzipReader, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()
	decompressedData, err := io.ReadAll(gzipReader)
	if err != nil {
		return nil, err
	}
	return decompressedData, nil
}
