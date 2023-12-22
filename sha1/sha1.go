package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println(sha1Sum("./http.log.gz"))
	fmt.Println(sha1Sum("./sha1.go"))

}

func sha1Sum(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("failed opening file: %w", err)
	}

	defer file.Close()

	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {
		gzipFile, err := gzip.NewReader(file)
		if err != nil {
			return "", fmt.Errorf("failed to open gzip reader: %w", err)
		}

		defer gzipFile.Close()

		r = gzipFile
	}

	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", fmt.Errorf("failed to copy to sha1: %w", err)
	}

	sig := w.Sum(nil)

	return fmt.Sprintf("%x", sig), nil
}
