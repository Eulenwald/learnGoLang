package main

import (
	"fmt"
	"math/rand"
)

/*
	Lektion 249 – Übung 7

Schreiben Sie ein Programm, das 10 Go-Routinen started und lassen Sie jede Go-Routine 10 Zahlen
in einen Channel schreiben. Empfangen Sie alle 100 Werte aus dem Channel.
*/
func uebung07() {
	fmt.Println("----------------")
	println("uebung07")

	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for n := 0; n < 10; n++ {				
				c <- rand.Int()
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(<-c)
	}
	/*
		for v := range c {
			fmt.Println(v)
		}
	*/
	fmt.Println("----------------")
}

func goRout(c chan int) {
	for i := 0; i < 10; i++ {
		i *= rand.Int()
		c <- i
	}
}
