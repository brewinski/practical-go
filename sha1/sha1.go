package main

import (
	"fmt"
	"os"
)

func main() {

}

func sha1Sum(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("failed opening file: %w", err)
	}

	defer file.Close()

	return "", nil
}
