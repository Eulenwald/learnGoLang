package main

import "fmt"

/*
	Lektion 241 – Übung 3

Starten Sie mit: https://go.dev/play/p/py-gC656Wjg und holen Sie die Werte mit Hilfe einer mit
„range“ verwendenden Schleife aus dem Channel (schließt Ausgabe ein).
*/
func uebung03() {
	fmt.Println("----------------")
	println("uebung03")

	c := gen()
	receive(c)

	fmt.Println("about to exit")
	fmt.Println("----------------")
}

// Sorgen Sie dafür, dass das
// Argument dieser Funktion ein
// "receive-only"-Channel ist
func receive(c <-chan int) {

	for v := range c {
		fmt.Println(v)
	}
	// Hier ist Platz für Ihren Code
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
