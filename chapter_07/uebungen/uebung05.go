// uebung05.go
/*
Lektion 56 – Übung 5
Erstellen Sie mit dem Short Declaration Operator eine Variable vom Typ string und weisen Sie
ihr als Wert mittels eines „raw string literal“ zu. Der Wert sollte einen Zeilenumbruch enthalten
OHNE einen „escaped character“ wie „\n“ zu nutzen.
*/
package main

import "fmt"

func main() {
	a := `Jonny loves Jenny:
	"But" Jenny loves Jon`
	fmt.Println(a)
}