package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	println("channellike")

	sl := []int{}
	slSize := 10

	for i := 0; i < slSize; i++ {
		sl = append(sl, rand.Intn(100))
	}

	ch := make(chan int)

	//go summary(sl, ch)

	portion := 2
	parts := slSize / portion
	i := 0
	//if parts * portion < slSize {parts += 1}
	for i < parts {
		// sl[0:(1*2)]
		// fmt.Printf("go summary(sl[%d:%d],ch)\n", i*portion, (i+1)*portion)
		go summary(sl[i*portion:(i+1)*portion], ch)
		i += 1
	}
	close(ch)
	// looking for the rest of slice for the modulo rest
	// if parts*portion < slSize {
		//fmt.Printf("1 more time: go summary(sl[%d:],ch)\n", i*portion)
		//go summary(sl[i*portion:], ch)	}	

	//fmt.Scanln()	

	chSum := 0

	for i := 0; i < parts; i++ {
		part := <-ch
		fmt.Println("part is:", part)
		chSum += part
	}

	fmt.Println("ch summe is:", chSum)
	fmt.Scanln()
	iterate()
}

func summary(sl []int, ch chan int) {
	summary := 0
	for _, v := range sl {
		summary += v
	}
	fmt.Println(sl, summary)
	ch <- summary
	//fmt.Println(<- ch)
	//close(ch)
}

func iterate() {

	ch := make(chan int)
	go fib(10, ch)

	for v := range ch {
		i := v
		fmt.Println(v, i)
		//part := <- v 		fmt.Println("part is:", part)		chSum += part
	}
}

func fib(n int, ch chan int) {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
		time.Sleep(1 * time.Second)
	}
	close(ch)
}
