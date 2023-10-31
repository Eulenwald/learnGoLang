// uebung001.go
/*
1. Erstellen Sie mit dem Short Declaration Operator die Variablen mit den Identifiern „x”, „y” und
„z” und weisen Sie ihnen folgenden Werte zu:
a) 23
b) „Papa Schlumpf”
c) true
2. Geben Sie die Werte der Variablen aus mit
a) einem einzelnen „print“-Statement
b) mehreren einzelnen „print“-Statement
*/
package main

import "fmt"

// Zugriffsbereich mgl. klein halten also nicht hier
// Ausserdem sind das kein short-operatoren
//var x = 23
//var y = "Papa Schlumpf"
//var z = true

func main() {
	// implizierte Typ-Zuweisung
	x := 23
	y := "Papa Schlumpf"
	z := true

	fmt.Println(x, y, z) // oder
	fmt.Printf("%v %v %v\n", x, y, z)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}
