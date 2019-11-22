package main

import "fmt"

func main() {

	val := 9
	p := &val
	fmt.Println("Wert: ", *p, ", Speicheradresse: ", p)

}
