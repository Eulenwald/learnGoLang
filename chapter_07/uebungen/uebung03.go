// uebung03.go
/*
Lektion 52 – Übung 3
Erstellen Sie sowohl „typed“ als auch „untyped“ Konstanten.
Geben Sie in einem Statement die zugewiesenen Werte und zugewiesenen/angenommenen Typen
aus.
*/
package main

import "fmt"

const (o1 = 10
 o2 int = 20)

func main() {

	fmt.Printf("%v %T\n%v %T\n",o1, o1, o2, o2)
}