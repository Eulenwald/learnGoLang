// uebung004.go
/*
Für diese Übung
1. Erstellen Sie Ihren eigenen Variablentyp basierend auf dem Typ int.
2. Erstellen Sie die Variable „x“ mit dem von Ihnen erstellen Typ mittels var
3. In der Funktion main()
a) Geben Sie den Wert von „x” aus
b) Geben Sie den Typ von „x” aus
c) Weisen Sie „x” den Wert 23 mit dem einfachen Zuweisungsoperator zu
d) Geben Sie nochmal den Wert von „x” aus
*/
package main

import "fmt"

type meinInt int

var x meinInt

func main() {
	// Ausgabe der Values
	fmt.Printf("Der wert von x ist ->%v<-\n", x)
	fmt.Printf("x ist vom Typ %T\n", x)
	x = 23
	fmt.Printf("Der wert von x ist ->%v<-\n", x)
}
