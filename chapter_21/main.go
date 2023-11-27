package main

import "fmt"

func main() {
	
	fmt.Println("Kapitel 21 Channels")

	biChan := make(chan string)
	biChan <- "Hallo Welt"

	println(biChan)
}