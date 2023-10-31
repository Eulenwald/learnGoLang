// uebung002.go
/*
1. Erstellen Sie mit dem Keyword var global die Variablen mit den Identifiern „x”, „y” und „z” und
deklarieren sie deren Typen als:
a) int
b) string
c) bool
2. Geben Sie die Werte der Variablen in der Funktion main() aus.
Zusatzfrage: Wie nennt man diese vom Compiler zugewiesenen Werte?
*/
package main

import "fmt"

//
var x int
var y string
var z bool

func main() {
	// Ausgabe der Zero-Values
	fmt.Printf("Der wert von x ist ->%v<-\n", x)
	fmt.Printf("Der wert von y ist ->%v<-\n", y)
	fmt.Printf("Der wert von z ist ->%v<-\n", z)
}
