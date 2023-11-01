// uebung04.go
/*
Lektion 54 – Übung 4
Schreiben Sie ein Programm, das
•einer Variable den Typ int und den Wert 23232 zuweist.
•Geben Sie diesen Wert nebeneinander als binär, dezimal und hexadezimal aus.
•Weisen Sie das um 1 nach links verschobene Bitmuster diesen Wertes einer neuen Variable
zu.
•Geben Sie den Wert dieser Variablen nebeneinander als binär, dezimal und hexadezimal aus.
Tipp: „Verb“ für die Ausgabe in Binärschreibweise ist „%b“ und Breite des Ausdrucks auf 32
Stellen erweistert und (nach links) mit Nullen aufgefüllt: „%032b“.
*/
package main

import "fmt"

var i int = 23232

func main() {
	fmt.Printf("%032b %d %#x\n", i, i, i)
	i2 := i << 1
	fmt.Printf("%032b %d %#x", i2, i2, i2)
}