//Übungen zu structs und
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	ch := make(chan int)

	go sum(ch)
	go sum(ch)
	go sum(ch)

	ch <- 15
	ch <- 10
	ch <- 6

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Printf("Time taken: %s\n", time.Since(start).String())
}

func sum(ch chan int) {
	sum := 0
	max := <-ch

	for i := 0; i < max; i++ {
		sum += i
		time.Sleep(100 * time.Millisecond)
	}

	ch <- sum
}
