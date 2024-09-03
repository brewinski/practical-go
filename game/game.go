package main

import "fmt"

type Item struct {
	X, Y int
}

func main() {
	var i1 Item

	fmt.Println(i1)
	fmt.Printf("i1: %#v \n", i1)

	i2 := Item{1, 2}

	fmt.Println(i2)
	fmt.Printf("i2: %#v \n", i2)
}
