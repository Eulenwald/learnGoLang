package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func sizes() {
	var myInt uint8
	var myFloat float64

	sizeInt1 := unsafe.Sizeof(myInt)
	sizeFloat1 := unsafe.Sizeof(myFloat)

	fmt.Printf("Size of int: %d bytes\n", sizeInt1)
	fmt.Printf("Size of float64: %d bytes\n", sizeFloat1)
	//	fmt.Fprintln(unsafe.Sizeof(myInt))

	sizeInt2 := int(reflect.TypeOf(myInt).Size())
	sizeFloat2 := int(reflect.TypeOf(myFloat).Size())

	fmt.Printf("Size of int: %d bytes\n", sizeInt2)
	fmt.Printf("Size of float64: %d bytes\n", sizeFloat2)

	sizeInt3 := reflect.TypeOf(myInt).Size()
	sizeFloat3 := reflect.TypeOf(myFloat).Size()

	fmt.Printf("Size of int: %d bytes\n", sizeInt3)
	fmt.Printf("Size of float64: %d bytes\n", sizeFloat3)
}
