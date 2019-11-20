package main

import (
	"fmt"
)
/*
	Das defer Schlüsselwort sorgt dafür, dass die danach folgenden Anweisungen auf den Stack gelegt und erst nach
	beenden der umschließenden Funktion via dem LIFO-Prinzip (Last in First out) ausgeführt werden. In diesem 
	Beispiel wird erst jedes Printf Statement auf den Stack gelegt und am Ende der Funktion beginnend mit dem  
	LETZTEN Printf Statment der Stack abgearbeitet. Dadurch wird die Ausgabe des übergebenen Arrays umgedreht.
*/
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
