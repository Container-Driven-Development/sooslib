package sooslib

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"

	concatenate "github.com/paulvollmer/go-concatenate"
)

// Hash files into single token
func Tokenizer(hashFiles []string) string {

	data, err := concatenate.FilesToBytes("\n", hashFiles...)

	check(err)

	h := sha1.New()
	_, copyErr := io.Copy(h, bytes.NewReader(data))

	check(copyErr)

	fileSum := fmt.Sprintf("%x", h.Sum(nil))

	return fileSum
}
