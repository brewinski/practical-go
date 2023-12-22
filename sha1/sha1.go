package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(sha1Sum("./http.log.gz"))
}

func sha1Sum(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("failed opening file: %w", err)
	}

	defer file.Close()

	r, err := gzip.NewReader(file)
	if err != nil {
		return "", fmt.Errorf("failed to open gzip reader: %w", err)
	}

	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", fmt.Errorf("failed to copy to sha1: %w", err)
	}

	sig := w.Sum(nil)

	return fmt.Sprintf("%x", sig), nil
}
