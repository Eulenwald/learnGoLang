package main

import "fmt"

func main() {

	fmt.Println("Kapitel 21 Channels")

	gerade := make(chan int)
	ungerade := make(chan int)
	ende := make(chan bool)

	// send
	go send(gerade, ungerade, ende)

	// channel schliessen Funktioniert hier NICHT!
	// you can't close linear
	//close(gerade)
	//close(ungerade)
	//close(ende)

	// recieve
	receive(gerade, ungerade, ende)
}

func send(g, u chan<- int, e chan<- bool) {
	for i := 1; i < 11; i++ {
		if i%2 == 0 {
			g <- i
		} else {
			u <- i
		}
	}
	close(g)
	close(u)
	e <- true
	//close(g)
	//close(u)
	close(e)
}

func receive(g, u <-chan int, e <-chan bool) {

	var slGerade []int
	var slUngerade []int
	var slEnd []bool

	for {
		select {
		case geradeZahl, ok := <-g:
			slGerade = append(slGerade, geradeZahl)
			fmt.Println("Gerade:", geradeZahl, ok)
			fmt.Println("slice slGerade:", slGerade)
		case ungeradeZahl, ok := <-u:
			slUngerade = append(slUngerade, ungeradeZahl)
			fmt.Println("Ungerade:", ungeradeZahl, ok)
			fmt.Println("slice slUngerade:", slUngerade)
		case ende, ok := <-e:
			slEnd = append(slEnd, ende)
			fmt.Println("Singnal Ende:", ende, ok)
			fmt.Println("slce slEnd:", slEnd)
			return
		}
	}

}

func receiveAlt(g, u <-chan int, e <-chan bool) {

	for {
		select {
		case v := <-g:
			fmt.Println("Gerade:", v)
		case v := <-u:
			fmt.Println("Ungerade:", v)
		case v, ok := <-e:
			fmt.Println("Singnal Ende:", v, ok)
			return
		}
	}

}
