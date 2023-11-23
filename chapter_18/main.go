package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	sizes()
	fmt.Println("kapitel 18")
	fmt.Fprintln(os.Stdout, "Oder?")
	fmt.Fprintln(os.Stderr, "Fehlermeldungen landen hier!")
	io.WriteString(os.Stderr, "das geht auch!\n")

	//pDec *= json.NewDecoder()
	sortieren()
	linkedList()
	sortieren2()

}
