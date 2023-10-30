package main

import "fmt"

func main() {
	fmt.Println("Hallo, hier startet unser kleines Programm!")

	foo()
	fmt.Println("Foo ist beendet, unser kleines Programm geht evt. weiter!")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i, i)
		}
	}

	bar()

}

func foo() {
	fmt.Println("Hier sind wir in foo()")
}

func bar() {
	fmt.Println("Hier sind wir in bar()")
}

// Sequences
// Loops / iterieren
// conditionals
