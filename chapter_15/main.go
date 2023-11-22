package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Kapitel 15\nUebung 1")
	uebung01()
	fmt.Println("\nUebung 2")
	uebung02()
	fmt.Println("\nUebung 3")
	uebung03()
	fmt.Println("\nUebung 4")
	uebung04()
	fmt.Println("\nUebung 5")
	uebung05()
	fmt.Println("\nUebung 6")
	uebung06()
	fmt.Println("\nUebung 7")
	uebung07()
	fmt.Println("\nUebung 8")
	uebung08()
	fmt.Println("\nUebung 9")
	uebung09()
	fmt.Println("\nUebung 10")
	uebung10()
}

/*
	Lektion 161 – Übung 1

In dieser Übung:
•Erstellen Sie eine Funktion mit dem Identifier „foo“, die einen Wert vom Typ int zurückgibt
•Erstellen Sie eine Funktion mit dem Identifier „bar“, die einen Wert vom Typ int und einen
einen Wert vom Typ string zurückgibt
•Erstellen Sie für beide funktionen einen gültigen Funktionsaufruf und weisen sie neu
erstellten Variablen die Rückgabewerte zu.
•Geben Sie die Werte aus.
*/
func foo1() int           { return 1 }
func bar1() (int, string) { return 2, "mit Sternchen" }
func uebung01() {
	iVal1 := foo1()
	iVal2, s := bar1()
	fmt.Println(iVal1, iVal2, s)
}

/*
	Lektion 163 – Übung 2

Erstellen Sie eine Funktion mit dem Identifier „foo“, die
•die einen variatischen Parameter vom Typ int entgegennimmt
•Übergeben Sie einen Wert vom Typ []int in adequater Weise an die Funktion
•Geben Sie eine Summe aller übergebenen Eingabewerte als Rückgabewert zurück
Erstellen Sie eine Funktion mit dem Identifier „bar“, die
•als Parameter Werte vom Typ []int entgegennimmt
•Geben Sie eine Summe aller übergebenen Eingabewerte als Rückgabewert (Typ int) zurück
*/
func foo2(sl ...int) (sum int) {

	for _, v := range sl {
		sum = sum + v
	}
	return
}
func bar2(sl []int) (sum int) {

	for _, v := range sl {
		sum = sum + v
	}
	return
}
func uebung02() {
	testSlic := []int{1, 2, 3, 4}
	iVal1 := foo2(testSlic...)
	iVal2 := bar2(testSlic)
	fmt.Println(iVal1, iVal2)
}

/*
	Lektion 165 – Übung 3

Erstellen Sie zwei Funktionen und rufen Sie sie auf. Sorgen Sie mit „defer“ dafür, dass der erste
Aufruf bis nach den zweiten Aufruf verzögert wird.
*/
func uebung03() {
	testSlic1 := []int{1, 2, 3, 4}
	testSlic2 := []int{1, 2, 3, 4, 5, 6}
	iVal1 := foo2(testSlic1...)
	iVal2 := bar2(testSlic2)
	defer fmt.Println(iVal1)
	fmt.Println(iVal2)
}

/*
	Lektion 167 – Übung 4

Erstellen Sie einen Typ mit unterliegendem Struct mit dem Identifier „person“
Wählen Sie angemessene Typen für die Elemente:
vorname
nachname
alter
•weisen Sie dem Typ „person“ eine Methode zu, die den Identifier „sagt“ hat.
Die Methode greift auf das in „person“ definierte Struct zu und gibt einen String mit Name
und Alter aus.
•Erstellen Sie einen Wert des Typs „person“.
•Rufen Sie die Methode für den Werte auf.
*/
type person struct {
	vorname  string
	nachname string
	alter    int
}

func (r person) sagt() {
	fmt.Printf("%s %s ist %d Jahre alt!", r.vorname, r.nachname, r.alter)
}
func uebung04() {
	student := person{
		vorname:  "Andreas",
		nachname: "Eulenwald",
		alter:    55,
	}
	student.sagt()
}

/*
	Lektion 169 – Übung 5

Erstellen Sie einen Typ „quadrat“ und einen Typ „kreis“ beruhend auf Structs.
Erstellen Sie eine Method „fläche“, die einen Wert vom Typ float64 zurückgibt und weisen Sie
beiden Typen die Methode zu.
Fläche eines Kreisen = π r 2
Fläche eines Quadrates = Seitenlänge * Seitenlänge
Erstellen Sie einen Typ „formen“, der ein Interface definiert, das durch die Implementierung der
Methode „fläche“ definiert ist.
Erstellen Sie eine Funktion mit dem Identifier „info“, die den Typ „formen“ entgegennimmt und die
Fläche ausgibt.
Erstellen Sie, bzw geben Sie aus mittels „info“:
Wert vom Typ „quadrat“.
Wert vom Typ „kreis“.
Rufen Sie die Funktion info für beide Werte auf!
*/
type forms interface {
	flaeche() float64
}

func info(f forms) {
	fmt.Println(f.flaeche())
}

type rectangle struct {
	width  float64
	length float64
}

func (r rectangle) flaeche() float64 {
	return r.length * r.width
}

type circle struct {
	radius float64
}

func (r circle) flaeche() float64 {
	return math.Pi * r.radius * 2
}

func uebung05() {
	r := rectangle {
		width: 10.0,
		length: 10.0,
	}	
	info(r)
	k := circle {
		radius: 2,
	}
	info(k)
}

/*
	Lektion 171 – Übung 6

Erstellen und nutzen Sie eine „anonyme“ Funktion.
*/
func uebung06() {
	f := func (i int) {
		fmt.Println(i * 2)
	} 
	f(2)
}

/*
	Lektion 173 – Übung 7

•Weisen Sie eine Variablen eine Funktion zu, (die irgendetwas tut) und rufen Sie diese
Funktion auf.
•Erweitern Sie das Beispiel so, dass die Funktion auch einen Wert als parameter
entgegennimmt und ausgibt.
*/
func uebung07() {
	fmt.Println("siehe uebung 6")
}

/*
	Lektion 175 – Übung 8

•Erstellen Sie eine Funktion, die eine Funktion als Rückgabewert liefert
•Weisen Sie die zurückgegebene Funktion einer Variablen zu
•Rufen Sie die Funktion durch die Variable auf.
*/

func test() func (i int) {
	return func (i int) {
		fmt.Println("I ist:", i)
	}
}
func uebung08() {
	f := test()
	f(42)
}

/*
	Lektion 177 – Übung 9

Erstellen Sie einen Callback, d.h. erstellen Sie eine Funktion, die eine Funktion (und ein einen
Funktionswert) als Parameter entgegennimmt. Dann übergeben Sie eine Funktion (die z.B. einen
String ausgibt) und einen Wert (zum Beispiel „You shall not pass!“) und bringen Sie zur
Ausführung.
*/
func myCallback(f func(string), s string) {
	f(s)
}
func uebung09() {
	f := func(s string) {
		fmt.Println(s)
	}
	myCallback(f, "You shall not pass!")
}

/*
	Lektion 179 – Übung 10

Closures „verkapseln“ den Geltungsbereich einer Variablen in einem Codeblock. Erstellen Sie eine
Funktion mit dem Identifier „undNochEinEis“, die innerhalb einer zurückgelieferten anonymen
Funktion eine Variable „soVieleEis“ hochzählt.
Erstellen Sie eine Variable „spongbob“ und eine Variable „patrick“, der Sie jeweils die Funktion
„undNochEinEis“ zurodnen. Rufen Sie die „spongbob“-Funktion drei mal auf, die „patrick“-
Funktion dreiundzwanzig mal.
*/
func uebung10() {}
