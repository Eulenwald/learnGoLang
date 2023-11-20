package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var input io.Reader
var (
	flagFile = flag.String("file", "", "file used for input")
	flagURL = flag.String("url", "", "auszulesende url")
)

func main() {

	flag.Parse()

	switch {
	case *flagFile != "":
		f, err := os.Open(*flagFile)
		if err != nil {
			fmt.Println("error opening teh file:", *flagFile, err)
			os.Exit(123)
		}
		defer f.Close()
		input = f
	case *flagURL != "":
		res, err := http.Get(*flagURL)
		if err != nil {
			fmt.Println("Page doesn't exist")
			os.Exit(234)		
		}
		// Hier gehe ich davon aus, dass Response geklappt hat
		defer res.Body.Close()
		input = res.Body


	default:
		input = os.Stdin

	}

	printMD5(input, os.Stdout)
}

func printMD5(r io.Reader, w io.Writer) {
	h := md5.New()
	if _, err := io.Copy(h, r); err != nil {
		log.Fatal("Error copy data to hash", err)
		os.Exit(1)
	}

	fmt.Fprintf(w, "%x", h.Sum(nil))

}

//  ..\gocat\gocat.exe go.mod
