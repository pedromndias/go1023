package main

import "fmt"

func main() {
	//Realitza un map de tres alumnes i tres notes.
	m := map[string]int{
		"Pedro": 8,
		"Joana": 7,
		"Patts": 5,
	}
	fmt.Println(m)
	fmt.Println(m["Pedro"])
}
