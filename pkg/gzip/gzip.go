package gzip

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
)

func Compress(data map[string]map[string]any) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)
	_, err = gzipWriter.Write(jsonData)
	if err != nil {
		return nil, err
	}
	err = gzipWriter.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func SingleCompress(data map[string]any) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)
	_, err = gzipWriter.Write(jsonData)
	if err != nil {
		return nil, err
	}
	err = gzipWriter.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decompress(compressedData []byte) (map[string]map[string]any, error) {
	buf := bytes.NewBuffer(compressedData)
	gzipReader, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()
	decompressedData, err := io.ReadAll(gzipReader)
	if err != nil {
		return nil, err
	}
	var data map[string]map[string]any
	err = json.Unmarshal(decompressedData, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func SingleDecompress(compressedData []byte) (map[string]any, error) {
	buf := bytes.NewBuffer(compressedData)
	gzipReader, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()
	decompressedData, err := io.ReadAll(gzipReader)
	if err != nil {
		return nil, err
	}
	var data map[string]any
	err = json.Unmarshal(decompressedData, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
