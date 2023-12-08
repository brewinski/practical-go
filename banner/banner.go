package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)
	banner("G☺", 6)
}

func banner(text string, width int) {
	padding := (width - utf8.RuneCountInString(text)) / 2

	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}

	fmt.Println(text)

	for i := 0; i < width; i++ {
		fmt.Print("-")
	}

	fmt.Println()
}
