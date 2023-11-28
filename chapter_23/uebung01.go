package main

import "fmt"

/*
	Lektion 237 – Übung 1

Bringen Sie diesen Codeschnipsel zum Laufen: https://go.dev/play/p/-DpZPo8o5JQ
a) mit Hilfe eines „func literal“ also einer „anonymen“ selbst-aufrufenden Funktion
oder alternativ (!):
b) indem Sie einen Buffer im Channel einsetzen.
*/
func uebung01() {
	fmt.Println("----------------")
	println("uebung01a")

	c1 := make(chan int)    // unbufferd
	c2 := make(chan int, 1) // bufferd

	go func() {
		c1 <- 42
	}()

	fmt.Println(<-c1)
	fmt.Println("----------------")
	println("uebung01b")
	c2 <- 23
	fmt.Println(<-c2)
	fmt.Println("----------------")
}
