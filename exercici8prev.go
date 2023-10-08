package main

import "fmt"

//Amb aquest exercici tens que realitzar una estructura condicional que representi un porter de discoteca que nomès deixa passar si ets major d'edat. Pero ara amb la possibilitat de donar una resposta per quan no pot passar.

func main() {
	edat := 19
	if edat >= 18 {
		fmt.Println("Pots accedir")
	} else {
		fmt.Println(("No pots passar"))
	}
}
