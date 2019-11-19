//Übung für Schleifen und Arrays
package main

import "fmt"

func main() {
	fmt.Println(fibonacci(10))
}

func fibonacci(max int) []int {
	/*	Warum nicht var zahlen [max]int?
		-> man kann auf diese Art keine Arrays erstellen, deren Länge zur Laufzeit berechnet wird
	*/

	zahlen := make([]int, max)

	zahlen[0] = 0
	zahlen[1] = 1

	for i := 2; i < max; i++ {
		zahlen[i] = zahlen[i-1] + zahlen[i-2]
	}

	return zahlen
}
