package main

import "fmt"

func main() {
	fmt.Println("Chapter 13 struct's")
	uebung01()
	fmt.Println("##########################")
	uebung02()
	fmt.Println("##########################")
	uebung03()
	fmt.Println("##########################")
	uebung04()
	fmt.Println("##########################")
}

/*
git init
git remote add origin <repository-url>

	Lektion 133  – Übung 1

Erstellen Sie Ihren eigenen Datentyp "person", der einen zugrundeliegenden Typ "struct" hat, so
dass die folgenden Daten speichern kann:
•Vorname
•Nachname
•Alter
•mehrere Lieblings-Eiscreme-Sorten
Erstellen Sie zwei Werte vom Typ person. Geben Sie die Werte mittels range aus, die sich in einem
Element vom Typ []string für die Lieblings-Eiscreme-Sorte angegeben sind.
*/
type person struct {
	Vorname                 string
	Nachname                string
	Alter                   int
	LieblingsEiscremeSorten []string
}

func uebung01() {
	andreas := person{
		Vorname:                 "Andreas",
		Nachname:                "Eulenwald",
		Alter:                   55,
		LieblingsEiscremeSorten: []string{"vanille", "schocko"},
	}
	denise := person{
		Vorname:                 "Denis",
		Nachname:                "Eulenwald",
		Alter:                   45,
		LieblingsEiscremeSorten: []string{"Erdnuss", "Erdbeere"},
	}

	fmt.Printf("%v mag ", andreas.Vorname)
	for _, v := range andreas.LieblingsEiscremeSorten {
		fmt.Printf("%v ", v)
	}

	fmt.Println()

	fmt.Printf("%v mag ", denise.Vorname)
	for _, v := range denise.LieblingsEiscremeSorten {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

/*
	Lektion 135  – Übung 2

Nehmen Sie den Code aus der vorherigen Übung und speichern Sie die Werte vom Typ person in
einer Map mit dem Schlüssel des Typs der den Nachnamen enthält. Greifen Sie auf jeden Wert in
der Map zu und geben Sie auch die Werte aus, die im Slice enthalten sind.
*/
func uebung02() {
	andreas := person{
		Vorname:                 "Andreas",
		Nachname:                "Eulenwald",
		Alter:                   55,
		LieblingsEiscremeSorten: []string{"vanille", "schocko"},
	}
	denise := person{
		Vorname:                 "Denis",
		Nachname:                "Eulenwald/Schroeder",
		Alter:                   45,
		LieblingsEiscremeSorten: []string{"Erdnuss", "Erdbeere"},
	}

	bewohner := make(map[string]person)
	bewohner[andreas.Nachname] = andreas
	bewohner[denise.Nachname] = denise

	for _, v := range bewohner {
		fmt.Println(v.Vorname)
		fmt.Println(v.Nachname)
		fmt.Println(v.Alter)
		fmt.Println(v.LieblingsEiscremeSorten)
	}

}

/*
	Lektion 137  – Übung 3

Erstellen Sie einen neuen Typ: fahrzeug.
•Der zugrunde liegende Typ ist ein struct.
•Die Felder:
- anzahlTüren
- farbe
•Erstellen Sie zwei neue Typen: lkw und pkw
•Der zugrunde liegende Typ jedes dieser neuen Typen ist ein struct.
•Betten Sie den Typ "fahrzeug" in lkw und pkw ein.
•Geben Sie dem lkw das Feld "vierrad", das auf bool gesetzt wird.
•Geben Sie der pkw das Feld "luxus", das auf bool gesetzt wird.
Verwendung der fahrzeug-, lkw- und pkw-Structs:
•Erstellen Sie mit einem Composite Literal „brummi“ und weisen Sie den Felder Werte zu.
•Erstellen Sie mithilfe eines Composite Literal „mittelklassewagen“ und weisen Sie den
Felder Werte zu.
•Geben Sie diese beiden „Values of Type“ aus.
•Geben Sie wenigstens einen der Werte der eingebetteten Felder aus.
*/
type fahrzeug struct {
	anzahlTueren int
	farbe        string
}
type lkw struct {
	fahrzeug
	vierrad bool
}
type pkw struct {
	fahrzeug
	luxus bool
}

func uebung03() {
	brummi := lkw{
		fahrzeug: fahrzeug{
			anzahlTueren: 2,
			farbe:        "schwarz",
		},
		vierrad: false,
	}
	mittelklasse := pkw{
		fahrzeug: fahrzeug{
			anzahlTueren: 4,
			farbe:        "blau",
		},
		luxus: true,
	}
	fmt.Println(brummi.anzahlTueren)
	fmt.Println(brummi.farbe)
	fmt.Println(brummi.vierrad)

	fmt.Println(mittelklasse.anzahlTueren)
	fmt.Println(mittelklasse.farbe)
	fmt.Println(mittelklasse.luxus)
}

/*
	Lektion 139  – Übung 4

Erstellen Sie ein „anonymous“ struct, also ein Struct ohne Identifier.
*/
func uebung04() {

	x:= struct {
		name string
	}{ name: "anoymes struct"}
	fmt.Println(x.name)
}
