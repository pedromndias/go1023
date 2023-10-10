package main

import "fmt"

func main() {
	notes := make([]int, 9, 10)

	notes = append(notes, 21)
	notes = append(notes, 22)

	fmt.Println(notes)
	fmt.Println(len(notes))
	fmt.Println(cap(notes))

}
