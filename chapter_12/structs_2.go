package main

import "fmt"

type schluempfe struct {
	vorname  string
	nachname string
	alter    uint8
}

type oberSchluempfe struct {
	schluempfe
	vorname     string
	kannZaubern bool
}

func main() {
	zauberschlumpf := oberSchluempfe{
		schluempfe: schluempfe{
			vorname:  "Papa",
			nachname: "Schlumpf",
			alter:    157,
		},
		vorname:     "Heinrich",
		kannZaubern: true,
	}
	schlumpf2 := schluempfe{
		vorname:  "Schlaumi",
		nachname: "Schlumpf",
		alter:    20,
	}

	fmt.Println(zauberschlumpf)
	fmt.Println(schlumpf2)

	fmt.Println(zauberschlumpf.schluempfe.vorname, zauberschlumpf.nachname, "kann zaubern:", zauberschlumpf.kannZaubern)
	fmt.Println(schlumpf2.vorname)
}
