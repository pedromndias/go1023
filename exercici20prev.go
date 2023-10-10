package main

import "fmt"

func main() {
	carro := []string{"pastanagues", "pebrots", "aigua", "formatge"}
	//Afegeix a el carro de la compra "Oli d'oliva"

	carro = append(carro, "Oli d'oliva")

	fmt.Println(carro)

	for i, v := range carro {
		fmt.Println(i, v)
	}

}
