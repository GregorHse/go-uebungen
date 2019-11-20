package main

import (
	"fmt"
)

func flipText(arr []rune) {
	fmt.Println("Der Text umgedreht lautet:")
	for i := 0; i < len(arr); i++ {
		defer fmt.Printf("%c", arr[i])
	}
}
func main() {
	var array = []rune{'d', 'l', 'r', 'o', 'w', ' ', 'o', 'l', 'l', 'e', 'h'}
	flipText(array)

}
