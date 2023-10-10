package main

import "fmt"

func main() {
	//A continuació realitza 2 structs anonims sobre articles de supermercats a on tens que definir les dades de nom, unitats i pvp.
	s1 := struct {
		nom     string
		unitats int
		pvp     float64
	}{
		nom:     "agua",
		unitats: 5,
		pvp:     1.3,
	}

	s2 := struct {
		nom     string
		unitats int
		pvp     float64
	}{
		nom:     "pan",
		unitats: 1,
		pvp:     2.6,
	}
	fmt.Println(s1)
	fmt.Println(s2)
}
