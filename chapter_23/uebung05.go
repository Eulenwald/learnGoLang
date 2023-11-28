package main

import "fmt"

/*
	Lektion 245 – Übung 5

Starten Sie mit: https://go.dev/play/p/7aWqxdYLgYJ und verwenden Sie zwei mal ein „,ok“
(Komma okay)-Statement (vor und nach dem close()), um zu zeigen, dass der Channel leer ist und
kein Wert mehr aus dem Channel stammt.
*/
func uebung05() {
	fmt.Println("----------------")
	println("uebung05")

	c := make(chan int)

	go func() {
		c <- 42
	}()

	// Platz für erstes ,ok-Statement
	v, ok := <-c
	fmt.Println("Ergebnisse:", v, ok)
	close(c)

	// Platz für zweites ,ok-Statement
	v, ok = <-c
	fmt.Println("Ergebnisse:", v, ok)

	fmt.Println("----------------")
}
