package utils

import (
	"archive/zip"
	"bytes"
)

type ZipWriter struct {
	zipWriter *zip.Writer
}

func NewZipWriter(buffer *bytes.Buffer) *ZipWriter {
	return &ZipWriter{zipWriter: zip.NewWriter(buffer)}
}

func (z *ZipWriter) AddFile(fileName string, data []byte) error {
	file, err := z.zipWriter.Create(fileName)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (z *ZipWriter) Close() error {
	return z.zipWriter.Close()
}
