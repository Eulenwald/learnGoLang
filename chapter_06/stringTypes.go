// stringTypes.go
/*
Lektion 42 – String ist ein Typ
Strings sind in Go ein eigener Datentyp.
*/
package main

import (
	"fmt"
)

func main() {
	aString := "Hallo ich bin ein String!"
	fmt.Println(aString)
	aString = "ABC, die katze liegt im Schnee!"
	fmt.Println(aString)

	aChar := aString[2]
	fmt.Printf("Value of aChar: %v\n", aChar)
	fmt.Printf("Type of aChar: %T\n", aChar)
	fmt.Printf("Char of aChar: %c\n", aChar)

	aSlice := []byte(aString)
	fmt.Printf("Value of aChar: %v\n", aSlice)
	fmt.Printf("Type of aChar: %T\n", aSlice)

	aString = "ABC世"
	aSlice = []byte(aString)
	fmt.Printf("Value of aChar: %v\n", aSlice)
	fmt.Printf("Type of aChar: %T\n", aSlice)

	aChar = aString[4]
	fmt.Printf("Value of aChar: %v\n", aChar)
	fmt.Printf("Type of aChar: %T\n", aChar)
	fmt.Printf("Char of aChar: %c\n", aChar)

	for i := 0; i < len(aString); i++ {
		fmt.Printf("%#U", aString[i])		
	}

	fmt.Println()
	for i, v := range aString {
		fmt.Printf("Index %d Hex: %#x\n", i, v)
	}

}
