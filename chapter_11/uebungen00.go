package main

import "fmt"

func main() {
	uebung08()
}

/*
	Lektion 107  – Übung 1

Benutzen Sie ein „composite literal“, um:
•ein Array mit 5 Elementen vom Typ int zu erzeugen
•Weisen Sie manuell jeder Indexposition einen Wert zu
•Nutzen Sie eine For-Schleife mit „range“ um das Array und Index auszugeben
•nutzen Sie eine formatierte Ausgabe
•und geben Sie anschließend den Typ des Arrays aus
*/
func uebung01() {
	a := [5]int{1, 2, 3, 4, 5}

	for i, v := range a {
		fmt.Printf("Index %v hat den Wert %v\n", i, v)
	}
	fmt.Printf("Das array ist vom Typ '%T'", a)
}

/*
	Lektion 109  – Übung 2

Benutzen Sie ein „composite literal“, um:
•erstellen Sie ein Slice aus Werten vom Typ int
•weisen Sie 10 Werte zu
•nutzen Sie eine For-Schleife mit „range“ um Werte des Slices und Index auszugeben
•nutzen Sie formatierte Ausgabe
•und geben Sie anschließend den Typ des Slice aus
*/
func uebung02() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range sl {
		fmt.Printf("Index %v hat den Wert %v\n", i, v)
	}
	fmt.Printf("Das slice ist vom Typ '%T'", sl)
}

/*
	Lektion 111  – Übung 3

Erstellen Sie folgendes Slice mit Werten vom Typ int:
[42 43 44 45 46 47 48 49 50 51]
Nutzen Sie „Slicing“ um folgende Ausgaben zu erreichen (ohne das Slice zu verändern)
[42 43 44 45 46]
[47 48 49 50 51]
[44 45 46 47 48]
[43 44 45 46 47 78 49 50]
*/
func uebung03() {
	sl := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(sl[:5])
	fmt.Println(sl[5:])
	fmt.Println(sl[2:7])
	sl[6] = 78
	fmt.Println(sl[1:9])
}

/*
	Lektion 113  – Übung 4

Führen Sie folgende Schritte aus.
Beginnen Sie mit folgendem Slice:
x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
•Hängen Sie mit append() den Wert 51 an
•Geben Sie das slice aus
•Hängen Sie ein einem Statement die Werte 52, 53 und 54 an
•Geben Sie das slice aus
•Hängen Sie mit einem Statement an das Slice das folgende Slice an
y := []int{56, 57, 58, 59, 60}
Geben Sie das Slice x aus
*/
func uebung04() {
	sl := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	sl = append(sl, 51)
	fmt.Println(sl)
	sl = append(sl, 52, 53, 54)
	fmt.Println(sl)

	y := []int{56, 57, 58, 59, 60}
	sl = append(sl, y...)
	fmt.Println(sl)
}

/*
	Lektion 115  – Übung 5

Führen Sie nachfolgende Schritte aus. Beginnen Sie mit folgendem Slice:
x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
Nutzen Sie append() und Slicing, um das folgende Slice dem neu zu erstellenden
Slice y zuzuweisen: [42, 43, 44, 48, 49, 50, 51]
*/
func uebung05() {
	sl := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(sl[:4], sl[6:]...)
	fmt.Println(y)
}

/*
	Lektion 117  – Übung 6

Erstellen Sie ein Slice, um die Namen aller deutschen Bundesländer zu speichern.
Verwenden Sie make() und append(), um dies zu tun.
Ziele:

	Das Array, welchess dem Slice zugrunde liegt, soll nicht mehr als einmal
	erstellt werden.
	Wie lang ist Ihr Slice? Wie groß ist die Kapazität?
	Geben Sie alle Werte zusammen mit mit ihrer Indexposition aus, ohne „range“ zu verwenden.

	(Aufgabe ist weniger einfach als sie aussieht.)

Hier die Bundesländer:
`Bayern`, `Baden-Württemberg`, `Berlin`, `Brandenburg`, `Bremen`,
`Hamburg`, `Hessen`, `Mecklenburg-Vorpommern`, `Niedersachsen`,
`Nordrhein-Westfalen`, `Rheinland-Pfalz`, `Saarland`, `Sachsen`,
`Sachsen-Anhalt`, `Schleswig-Holstein`, `Thüringen`
*/
func uebung06() {
	bl := make([]string, 0, 16)
	bl = append(bl, `Bayern`, `Baden-Württemberg`, `Berlin`, `Brandenburg`, `Bremen`,
		"Hamburg", `Hessen`, `Mecklenburg-Vorpommern`, `Niedersachsen`,
		`Nordrhein-Westfalen`, `Rheinland-Pfalz`, `Saarland`, `Sachsen`,
		`Sachsen-Anhalt`, `Schleswig-Holstein`, `Thüringen`)

	fmt.Println(bl)
	for i := 0; i < 16; i++ {
		fmt.Printf("An Index %2d befindet sich %s\n", i, bl[i])
	}
	fmt.Printf("Der slice ist %v lang und hat eine Kapazitaet von %d", len(bl), cap(bl))
}

/*
	Lektion 119  – Übung 7

Erstellen Sie ein Slice von Slice von String ([][]string). Speichern Sie die
folgenden Werte:

	"James", "Bond", "Bond, James Bond"
	"Papa", "Schlumpf", "Schlumpf, Papa Schlumpf"
	"Rick", "Sanchez", "Schlauster Kopf im Multiversum"
	"Morty", "Smith", "Hauptberuflicher Sidekick"

Nutzen Sie Sie zwei ineinander verschachtelte For-Schleifen mit range die Nummer
(Index) des Slices auszugeben und jeweils darunter alle Werte des jeweiligen
Slice mit vorangestellter Position innerhalb des jeweiligen Slice.
In Etwa:
Slice Nummer: 0
Postion 0: James
Postion 1: Bond
Postion 2: Bond, James Bond
usw ...
*/
func uebung07() {
	slsl := [][]string{
		{"James", "Bond", "Bond, James Bond"},
		{"Papa", "Schlumpf", "Schlumpf, Papa Schlumpf"},
		{"Rick", "Sanchez", "Schlauster Kopf im Multiversum"},
		{"Morty", "Smith", "Hauptberuflicher Sidekick"},
	}

	for i, v := range slsl {
		fmt.Printf("Slice Number: %d\n", i)
		for i, v := range v {
			fmt.Printf("\tPosition %d: %s\n", i, v)
		}
	}
}

/*
	Lektion 121  – Übung 8

Erstellen Sie eine map mit einem Key vom Typ string, der "Vorname Familienname" einer Person
entspricht, und einem Werten vom Typ []string, die ihre Lieblingsdinge speichert. Speichern Sie
sieben Datensätze in Ihrer map. Geben Sie alle Werte aus, zusammen mit ihrer Indexposition im
Slice aus.
`Stan Smith`, `Amerika`, `Familie`, `Jesus`
`Francine Smith`, `Lippenstift`, `Pinke Kleider`, `Weinen unter der Dusche`
`Hayley Smith`, `Stirnband`, `Tank Top`, `Sandalen`
`Steve Smith`, `Computer`, `Mädchen`, `Freunde`
`Roger Smith`, `Fernsehen`, `Alkohol`, `Drogen`
`Klaus Heissler`, `Skispringen`, `Schwimmen`, `Rap & Hip Hop`
`Jeff Fischer`, `Gras rauchen`, `Fish (die Band)`, `seinen Hut`
*/
func uebung08() {

	m := map[string][]string{
		`Stan Smith`:     {`Amerika`, `Familie`, `Jesus`},
		`Francine Smith`: {`Lippenstift`, `Pinke Kleider`, `Weinen unter der Dusche`},
		`Hayley Smith`:   {`Stirnband`, `Tank Top`, `Sandalen`},
		`Steve Smith`:    {`Computer`, `Mädchen`, `Freunde`},
		`Roger Smith`:    {`Fernsehen`, `Alkohol`, `Drogen`},
		`Klaus Heissler`: {`Skispringen`, `Schwimmen`, `Rap & Hip Hop`},
		`Jeff Fischer`:   {`Gras rauchen`, `Fish (die Band)`, `seinen Hut`},
	}
	

	for k, v := range m {
		fmt.Printf("%v hat als Hobbies:\n", k)
		for i, stringAusSlice := range v {
			fmt.Printf("\t%v %v\n", i, stringAusSlice)
		}
	}
}

/*
*/
func uebung09(){

}
