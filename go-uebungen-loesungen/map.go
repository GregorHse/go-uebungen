//Übung zu maps
package main

import "fmt"

func main() {
	var words = []string{"go", "lang", "go", "gopher"}
	fmt.Println(countWords(words))

}

func countWords(words []string) map[string]int {
	var m = make(map[string]int)

	for _, word := range words {
		m[word]++
	}

	return m
}
