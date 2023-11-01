// uebung01.go
/*
Lektion 48 – Übung 1
Schreiben Sie ein kurzes Programm, das einer Variablen den Typ uint32 zuweist. Weisen Sie dieser
Variablen einen Wert aus ihrem Wertebereich zu.
Geben Sie den Wert als Dezimalzahl, als Binärzahl und in hexadezimaler Schreibweise (mit
vorangestelltem „0x“) aus.
*/
package main

import "fmt"

var iTest uint32 = 42

func main() {
	fmt.Printf("%d %b %#x", iTest, iTest, iTest)
}