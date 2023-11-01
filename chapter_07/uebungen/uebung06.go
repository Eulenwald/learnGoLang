// uebung06.go
/*
Lektion 58 – Übung 6
Erstellen Sie 4 Konstanten der nächsten Jahreszahlen beginnenend mit dem aktuellen Jahr und
nutzen Sie dazu den Ausdruck iota in allen Wertzuweisungen. Geben Sie die Konstanten
nebeneinander durch Leerzeichen getrennt aus.
*/
package main

import "fmt"

const (
	jahr = 2023 + iota // iota == 0
	jahr1 = jahr + iota // iota == 1
	jahr2 = jahr + iota // iota == 2
	jahr3 = jahr + iota // iota == 3
)

func main() {
	fmt.Println(jahr, jahr1, jahr2, jahr3)
}
