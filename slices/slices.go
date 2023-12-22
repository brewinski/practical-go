package main

import "fmt"

func main() {
	var s []int

	fmt.Println(len(s))
	fmt.Printf("s = %#v \n", s)

	s2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("s2 = %#v \n", s2)

	s3 := s2[1:4]
	fmt.Printf("s3 = %#v \n", s3)

	s3 = append(s3, 100)
	fmt.Printf("s3 append = %#v \n", s3)
	fmt.Printf("s2 append = %#v \n", s2)
}
