package main

import (
	"fmt"
	"time"
)

func main() {
	uebung06()
	testSprung()
}

func uebung01(){
	for i := 1; i < 10001; i++ {
		fmt.Println(i)
	}
}

func uebung02(){
	for i := 65; i < 91; i++ {
		fmt.Printf("%c\n", i)

		for x:=0; x < 3; x++ {
			fmt.Printf("%#U\n", i)
		}
	}
}

func uebung03(){
	var birthdayYear int = 1968

	// Aktuelles Datum und Uhrzeit abrufen
	aktuellesDatum := time.Now()
	// Das aktuelle Jahr extrahieren
	aktuellesJahr := aktuellesDatum.Year()
	
	for i := birthdayYear; i <= aktuellesJahr; i++ {
		fmt.Printf("%v\n", i)
	}
}

func uebung04(){
	var birthdayYear int = 1968
	// same like uebung03 but 1 step
	var actYear int = time.Now().Year()
	for i := birthdayYear; ; i++ {
		fmt.Printf("%v\n", i)
		if i >= actYear {break}
	}
}

/**
Lektion 82 – Übung 5
Geben Sie den Rest (modulo) aus, der bei der Teilung der Zahlen zwischen 10 und 100 (jeweils 
einschließlich) durch 4 ergibt.
*/
func uebung05(){
	for i := 8; i <= 100; i++ {
		fmt.Printf("Der Rest von '%3d %% 4' ist: %d\n", i, (i%4))
	}
}

/**
Lektion 84 – Übung 6
Erstellen Sie ein Programm, das ein if-Statement nutzt.
*/
func uebung06(){
	if 'a' < 'b' {
		fmt.Println("a ist < b")
	}
}


func testSprung() {

	sprungmarke:
	for i := 0; i < 20; i++ {
		if i == 5 { break sprungmarke}
		fmt.Println(i)
	}
}