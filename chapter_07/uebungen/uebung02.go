// uebung02.go
/*
Lektion 50 – Übung 2 
Nutzen Sie die nachfolgenden Operatoren und erstellen Sie je einen Ausdruck damit, deren Werte 
(Auswertung) Sie jeweils einer Variablen zuweisen (mittels Short Declaration Operator).
1. == in Variable a
2. <= in Variable b
3. >= in Variable c
4. != in Variable d
5. < in Variable e
6. > in Variable f
Geben Sie die Werte der Variablen a bis f mit nur einem Statement untereinander aus.
*/
package main

import "fmt"



func main() {
	o1 := 10
	o2 := 20
	a := o1 == o2
	b := o1 <= o2
	c := o1 >= o2
	d := o1 != o2
	e := 01 < o2
	f := o1 < o2

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n",a, b, c, d, e, f)
}