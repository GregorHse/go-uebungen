//Übung zu Slices und Ranges
package main

import "fmt"

func main() {
	var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	print(numbers[:3])
	print(numbers[len(numbers)-3:])
	print(numbers[3:7])

}

func print(numbers []int) {
	for _, value := range numbers {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
}
