package main

import (
	"cmp"
	"fmt"
	"slices"
)

func sortieren3() {

	p1 := person{"Peter", 32}
	p2 := person{"Bob", 34}
	p3 := person{"Justus", 31}
	p4 := person{"Martha", 56}
	p5 := person{"Schlumpfine", 23}
	p6 := person{"Gargamel", 64}
	p7 := person{"Rick", 75}
	p8 := person{"Morty", 13}
	p9 := person{"Summer", 17}

	people := []person{p1, p2, p3, p4, p5, p6, p7, p8, p9}

	fmt.Println(people)
	// Anonym
	slices.SortFunc(people, func(a, b person) int {
		if n := cmp.Compare(a.name, b.name); n != 0 {
			return n
		}
		// If names are equal, order by age
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(people)

	// Sort by Name
	slices.SortFunc(people, sortByName)
	fmt.Println(people)

	// Sort by Age
	slices.SortFunc(people, sortByAge)
	fmt.Println(people)

}

func sortByName(a, b person) int {
	if n := cmp.Compare(a.name, b.name); n != 0 {
		return n
	}
	// If names are equal, order by age
	return cmp.Compare(a.age, b.age)
}

func sortByAge(a, b person) int {
	return cmp.Compare(a.age, b.age)
}
