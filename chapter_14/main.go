package main

import "fmt"

func main() {
	fertigeSumme := summe(1, 2, 3, 4, 5, 6)
	fmt.Printf("Die summer aller Seiten ist %v", fertigeSumme)
}

func summe(x ...int) int {

	sum := 0

	for _, v := range x {
		fmt.Printf("%v + %v ist = ", sum, v)
		sum += int(v)
		fmt.Printf("%v.\n", sum)
	}

	return sum

}