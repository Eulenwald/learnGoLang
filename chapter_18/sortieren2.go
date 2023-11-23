package main

import (
	"fmt"
	"sort"
)

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

type ByName []person

func (a ByName) Len() int      { return len(a) }
func (a ByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool {
	if a[i].name == a[j].name {
		return a[i].name < a[j].name
	} else {
		return a[i].name < a[j].name
	}
}

func sortieren2() {

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
	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)
}
