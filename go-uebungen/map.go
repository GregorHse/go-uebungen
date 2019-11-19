//Übung zu maps
package main

import "fmt"

func main() {
	var words = []string{"go", "lang", "go", "gopher"}
	fmt.Println(countWords(words))
}

/*
	Schreiben sie diese Methode so, dass eine map zurückgegeben wird,
	die die Wörter und ihre häufigkeit enthält. Obiger Aufruf sollte
	map[go:2 gopher:1 lang:1]
	auf die Konsole drucken
*/
func countWords(words []string) map[string]int {

}
