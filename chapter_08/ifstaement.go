package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界 loop")
	// 32 - 126
	

	// Ausgabe Hexadezimal, Dezimal(ASCII) zeichen und Unicode-Zeichen
	for i := 32; i < 127; i++ {
		fmt.Printf("Hex %#x bzw. der ASCII-Code %03d entspricht dem Zeichen %c\t%#U\n",i,i,i,i)
	}
}