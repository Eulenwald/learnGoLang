package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	//os.arg []string startparameter
	fmt.Println(os.Args[1:])

	if(len(os.Args) < 2) {
		log.Println("Error: minimum 1 args missed")
		os.Exit(1)
	}

	// loop over the args
	for _, arg := range os.Args[1:]{

		pf, err := os.Open(arg)
		if(err != nil) {
			log.Println("Oh Oh something shit is happen. can not open file:", arg)
			os.Exit(2)
		}
		_, err = io.Copy(os.Stdout, pf)
		if(err != nil){
			log.Println("erro by io.Copy")
		}
		pf.Close()
		fmt.Println("read file:", arg)
	}
}

// func Copy(dst Writer, src Reader) (written int64, err error)