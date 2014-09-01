package util

import (
	"bytes"
	"compress/zlib"
	"io"
)

func CompressString(s string) (string, error) {
	var b bytes.Buffer
	compress := zlib.NewWriter(&b)

	_, err := compress.Write([]byte(s))
	if err != nil {
		return b.String(), err
	}

	compress.Close()

	return b.String(), nil
}

func UncompressString(s string) (string, error) {
	var b bytes.Buffer
	str_buffer := bytes.NewBufferString(s)

	uncompressed, err := zlib.NewReader(str_buffer)
	if err != nil {
		return "", err
	}

	io.Copy(&b, uncompressed)
	uncompressed.Close()

	return b.String(), nil
}
