// type type type type type
// casting in go ist conversion
package main

import "fmt"

var a int
type ganzzahl int
var b ganzzahl

func main() {
	a = 23
	fmt.Printf("a hat den Wert: %v\n",a)
	fmt.Printf("a hat den Wert: %b\n",a) // 10111
	fmt.Printf("a ist vom Typ: %T", a)


	b = 42
	fmt.Printf("b hat den Wert: %v\n",b)
	fmt.Printf("b ist vom Typ: %T", b)

	// Das funktioniert nicht
	// a = b
	a = int(b)
	fmt.Printf("a hat den Wert: %v\n",a)

}