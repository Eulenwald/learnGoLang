package main

import "fmt"

type emailaddress struct {
	name string
}

type schluempfe struct {
	nickname string
	name     string
	age      uint8
	email    emailaddress
}

func main() {
	schlumpf1 := schluempfe{
		nickname: "Papa",
		name:     "Schlupmf",
		age:      107,
		email:    emailaddress{name: "papa@web.de"},
	}

	schlumpf2 := schluempfe{
		nickname: "Schlaumi",
		name:     "Schlupmf",
		age:      20,
		email:    emailaddress{name: "schlaumi@web.de"},
	}

	fmt.Println(schlumpf1)
	fmt.Println(schlumpf2.email.name)
}
