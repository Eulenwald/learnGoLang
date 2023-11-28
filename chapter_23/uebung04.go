package main

import "fmt"

/*
	Lektion 243 – Übung 4

Starten Sie mit: https://go.dev/play/p/YxHwstTc3Jc und holen Sie die Werte mit Hilfe einer mit
„select“-Statement aus dem Channel (schließt Ausgabe ein).
*/
func uebung04() {
	fmt.Println("----------------")
	println("uebung04")

	q := make(chan int)
	c := gen4(q)

	receive4(c, q)

	fmt.Println("about to exit")

	fmt.Println("----------------")
}

// Sorgen Sie dafür, dass das
// Argument dieser Funktion ein
// "receive-only"-Channel ist
func receive4(c, q <-chan int) {

	// Hier ist Platz für Ihren Code
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q: // v := <-q:
			//fmt.Println(v)
			return
		}
	}
}

func gen4(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}
