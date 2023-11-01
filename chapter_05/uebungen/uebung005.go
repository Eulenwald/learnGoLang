// uebung005.go
/*
Auf Grundlage ihres Codebeispieles aus der vorherigen Übung
1. Erstellen mit var eine Variable mit dem Identifier „y“ auf der Gültigkeitsbereichebene des ganzen
Paketes und weisen Sie ihr den „underlying type“ ihres eigenen erstellen Types zu (also int).
2. In Funktion main()
Weisen Sie „y“ den Wert von „x“ zu und nutzen Sie „conversion“, also die Umwandlung des Wertes
der Variablen „x“ in den underlying type int.
3. Geben Sie den Wert von „y“ aus
4. Geben Sie den Typ von „y“ aus
*/
package main

import "fmt"

type meinInt int

var x meinInt
var y int

func main() {
	// Conversion meinInt -> int
	x = 23
	fmt.Printf("Der wert von x ist ->%v<-\n", x)
	fmt.Printf("x ist vom Typ %T\n", x)

	y = int(x)
	fmt.Printf("Der wert von y ist ->%v<-\n", y)
	fmt.Printf("y ist vom Typ %T\n", y)
	
}
