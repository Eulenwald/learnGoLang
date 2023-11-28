package main

import (
	"fmt"
)

/*
	Lektion 247 – Übung 6

Schreiben Sie ein Programm, das 100 Werte in einen Channel schreibt (an einen Channel sendet)
und anschließend alle Werte von diesem Channel empfängt und ausgibt.
*/
func uebung06() {
	fmt.Println("----------------")
	println("uebung06")

	ch := make(chan int)
	// var wg sync.WaitGroup

	go func() {
		//wg.Add(1)
		for i := 0; i < 100; i++ {
			ch <- i
		}
		//wg.Done()
		close(ch)
	}()
	//wg.Wait()
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("----------------")
}

// WaitGroup isn't nessessary
