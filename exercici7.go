package main

import "fmt"

func main() {
	edat := 28
	dia := "dimecres"
	passVip := true

	// Estructura condicional
	if (edat >= 18) || (dia == "dimecres" && passVip) {
		fmt.Println("Pots passar")
	} else if edat == 0 {
		fmt.Println("Tens que dir una edat correcte")
	} else {
		fmt.Println("No pots passar")
	}
}
