package main

import (
	"fmt"
	"time"
)

func say(s string, i int) {
	for index := 0; index < i; index++ {
		time.Sleep(100 * time.Millisecond)
		println(s)
	}
}

func main() {
	println("threads in go")

	go say("Hello", 3)
	go say("World", 2)

	fmt.Scanln()

	biChan := make(chan string)
	go sendData(biChan)
	go receiveData(biChan)
	fmt.Scanln()
}

// The function blocks as long the unbuffered channel has a value.
func sendData(c chan string) {
	println("string will send to channel")

	//time.Sleep(2 * time.Second)

	c <- "Hallo"
	// just here the unbuffered channel blocks the go-function as long the
	// channel has a value inside

	println("the string is rereceive from channel")
}

// The function blocks until it finds a value inside the unbuffered channel.
func receiveData(c chan string) {
	time.Sleep(2 * time.Second)
	println("string will receive from channel", <-c)
}
