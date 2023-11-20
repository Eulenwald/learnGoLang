package main

import "fmt"

type emailaddress struct {
	name string
}

type schluempfeA struct {
	nickname string
	name     string
	age      uint8
	email    emailaddress
}

func main2() {
	schlumpf1 := schluempfeA{
		nickname: "Papa",
		name:     "Schlupmf",
		age:      107,
		email:    emailaddress{name: "papa@web.de"},
	}

	schlumpf2 := schluempfeA{
		nickname: "Schlaumi",
		name:     "Schlupmf",
		age:      20,
		email:    emailaddress{name: "schlaumi@web.de"},
	}

	fmt.Println(schlumpf1)
	fmt.Println(schlumpf2.email.name)
}
