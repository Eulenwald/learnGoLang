package main

import (
	"fmt"
	"sort"
)

func sortieren() {

	sliceOfInts := []int{12345, 10, 8, 4, -1, 15, 88, 23, 43, 666}
	sliceOfStrings := []string{"Peter", "Bob", "Justus", "Charlie", "Mary", "Alice", "Bobby", "Zebra", "Aal", "a3", "a2"}

	fmt.Println(sliceOfInts)
	sort.Ints(sliceOfInts)
	fmt.Println(sliceOfInts)
	fmt.Println()
	fmt.Println(sliceOfStrings)
	sort.Strings(sliceOfStrings)
	fmt.Println(sliceOfStrings)
}
