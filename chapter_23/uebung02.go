package main

import "fmt"

/*
	Lektion 239 – Übung 2

Bringen Sie diese Codeschnipsel zum Laufen:
a) https://go.dev/play/p/_DBRueImEq
b) https://go.dev/play/p/oB-p3KMiH6
*/
func uebung02() {
	fmt.Println("----------------")
	println("uebung03")

	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)

	fmt.Println("----------------")
	println("uebung02b")

	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)

	fmt.Println("----------------")
}
