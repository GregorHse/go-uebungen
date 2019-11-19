//Übungen zu Goroutinen
package main

import (
	"fmt"
	"time"
)

/*
	Bauen Sie diese Methode so um, dass die Summen
	parallel statt nacheinander berechnet werden.
	Nutzen Sie Goroutinen und Channels.
	Die sum()-Methode muss ebenfalls verändert werden.
*/
func main() {
	start := time.Now()

	a := sum(15)
	b := sum(10)
	c := sum(6)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("Time taken: %s\n", time.Since(start).String())
}

func sum(max int) int {
	sum := 0

	for i := 0; i < max; i++ {
		sum += i
		time.Sleep(100 * time.Millisecond)
	}

	return sum
}
