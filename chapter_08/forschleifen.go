package main

import "fmt"

func main() {

	i := 0
	fmt.Println("Hello, 世界 1")
	for {

		fmt.Printf("i hat den Wert %v\n", i)
		if i > 9 {
			break
		}
		i++
	}
	fmt.Println("Ende der Welt 1")

	i = 1
	fmt.Println("Hello, 世界 2")
	for {
		if i > 9 {
			break
		}
		if i%2 == 0 {
			fmt.Printf("i hat den Wert %v\n", i)
		}
		i++
	}
	fmt.Println("Ende der Welt 2")

	i = 0

	fmt.Println("Hello, 世界 3")
	for {
		i++
		
		if i > 9 {
			break
		}
		if i%2 == 0 {
			continue
		}
		fmt.Printf("i hat den Wert %v\n", i)

	}
	fmt.Println("Ende der Welt 3")
}
