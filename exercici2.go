package main

import "fmt"

func main() {

	var b string
	var c float64
	var d bool

	// Verbs only work with Printf, not Println
	fmt.Printf("var a %T = %+v\n", "test", 6) //var a int =  0
	fmt.Printf("var b %T = %q\n", b, b)       //var b string = ""
	fmt.Printf("var c %T = %+v\n", c, c)      //var c float64 = 0
	fmt.Printf("var d %T = %+v\n\n", d, d)    //var d bool = false
}
