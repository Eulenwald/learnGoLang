package main

import (
	"fmt"
	"log"
	"math"
)

	type humans struct {
		nickname string
		name     string
		age      int
	}

	type employees struct {
		nickname string
		name     string
		age      int
		salary   int
		job      string
	}

	type externEmployees struct {
		nickname string
		name     string
		age      int
		binding  bool
		honorar  int
		job      string
	}

	type allEmployees interface {
		showEmployee()
	}

type rectangle struct {
	width  float32
	length float32
}

type circle struct {
	radius float32
}

type geometry interface {
	showArea() float32
	showPerimeter() float32
}

func (r rectangle) showArea() (rVal float32) {
	rVal = r.length * r.width
	fmt.Println("Die Flaeche des Rechtecks betraegt:", rVal)
	return
}
func (r rectangle) showPerimeter() (rVal float32) {
	rVal = 2 * (r.length + r.width)
	fmt.Println("Der Umfang des Rechtecks betraegt:", rVal)
	return
}

func (r circle) showArea() (rVal float32) {
	rVal = math.Pi * r.radius * r.radius
	fmt.Println("Die Flaeche des Kreises betraegt:", rVal)
	return
}
func (r circle) showPerimeter() (rVal float32) {
	rVal = 2 * r.radius * math.Pi
	fmt.Println("Der Umfang des Kreises betraegt:", rVal)
	return
}

func mesure(g geometry) {
	fmt.Println(g)
	g.showArea()
	g.showPerimeter()
}

func main() {
	/* 143
	fertigeSumme := summe(1, 2, 3, 4, 5, 6)
	fmt.Printf("Die summer aller Seiten ist %v", fertigeSumme) */

	/* 144
	xi := []uint8{0, 1, 1, 2, 3, 5, 8, 13, 21}
	fertigeSumme := summe(xi...)
	fmt.Printf("Die summer aller Seiten ist %v", fertigeSumme) */

	/* 145 defer
	defer printOut("Aufruf 1\n")
	printOut("Aufruf 3\n") */

	/* 146 Methods
	andreas := employees{
		nickname: "Andreas",
		name:   "der manhafte",
		age:    55,
		job:    "Junior Developer",
		salary: 42000,
	}
	frank := employees{
		nickname: "Frank",
		name:   "Schoeneshaar",
		age:    56,
		salary: 48000,
		job:    "Senior Developer",
	}

	peter := externEmployees{
		nickname: "Peter",
		name:     "Lustig",
		binding:  true,
		age:      80,
		honorar:  0,
		job:      "externer Fanboy",
	}

	//andreas.showEmployee()
	showMotto(andreas)
	//frank.showEmployee()
	showMotto(frank)
	//peter.showEmployee()
	showMotto(peter)
	*/

	/* 152 */
	r := rectangle{
		width:  5,
		length: 10,
	}
	k := circle{
		radius: 5,
	}

	mesure(r)
	mesure(k)
	
}

// 150
func (r externEmployees) showEmployee() {
	fmt.Println(r.nickname, r.name, r.age, r.binding)
	fmt.Println(r.job, r.honorar)
	fmt.Println()
}

func showMotto(aag allEmployees) {
	aag.showEmployee()
	switch aag.(type) {
	case employees:
		if aag.(employees).nickname == "Andreas" {
			fmt.Println("Alles nur gelogen")
		}
		fmt.Println("Ich arbeite gern fuer meinen Konzern")
	case externEmployees:
		fmt.Println("Ich arbeite gern mit dieser Firma")
	}
	fmt.Println()
}

// 146
/* Durch den receiver wird die Funktion yur Methode von human */
func (r employees) showEmployee() {
	fmt.Println(r.name, r.age)
	fmt.Println(r.job)
	fmt.Println(r.salary)
	fmt.Println()
}

/*
	Durch den PointerToReceiver sind Aenderungen  am zugrundeliegen

Objekt mgl
*/
func (r *employees) setAge(i int) {
	if i > 0 {
		log.Printf("changed age from %d to %d", r.age, i)
		r.age = i
	}
}

// 145 use defer
func printOut(s string) {
	fmt.Print(s)
}

// 144 + 143 variatische Parameter
func summe(x ...uint8) (sum int) {
	//sum := 0

	for _, v := range x {
		fmt.Printf("%v + %v ist = ", sum, v)
		sum += int(v)
		fmt.Printf("%v.\n", sum)
	}

	return
}
