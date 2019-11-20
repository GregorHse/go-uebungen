package main

import (
	"fmt"
)

/*
	Schreibe folgende Funktion so, dass das übergebene rune Array in
	umgekehrter Reihenfolge ausgegeben wird. Nutze dafür das defer Schlüsselwort.
	
	Tipp: Da es sich hier um Werte des Typ rune handelt, ist die Printf Funktion 
	hilfreich um die einzelnen runes als chars auszugeben.

*/

func flipText(arr []rune) {
	fmt.Println("Der Text umgedreht lautet:")

}
func main() {
	var array = []rune{'d', 'l', 'r', 'o', 'w', ' ', 'o', 'l', 'l', 'e', 'h'}
	flipText(array)

}
