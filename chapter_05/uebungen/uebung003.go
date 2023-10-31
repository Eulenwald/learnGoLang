// uebung003.go
/*
Auf Grundlage ihres Codebeispieles aus der vorherigen Übung
1. weisen Sie den drei Variablen auf der Gültigkeitsbereichebene des ganzen Paketes die Werte
a) 23
b) „Schlumpfine”
c) true
zu und in Funktion main()
2. Benutzen Sie die Funktion main Sprintf
a) um alle drei Werte, einer eigenen Variablen mit dem Identifier „s“ zuzuweisen, die Sie
mittels Short Declaration Operator erstellen,
b) und geben Sie den in „s“ gespeicherten Wert aus.
*/
package main

import "fmt"

var x int = 23
var y string ="Schlumpfine"
var z bool = true

/* Geht hier nicht muss bereis beim deklarieren passieren
  x = 23
	y = "Schlumpfine"
	z = true
*/

func main() {
	// Ausgabe der Values
	fmt.Printf("Der wert von x ist ->%v<-\n", x)
	fmt.Printf("Der wert von y ist ->%v<-\n", y)
	fmt.Printf("Der wert von z ist ->%v<-\n", z)

	
	sVal := fmt.Sprintf("Ist %s aelter als %d? %v", y, x, z)
	fmt.Println(sVal)

}
