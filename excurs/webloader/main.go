package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

//
var (
	pFlagOutput *string = flag.String("o", "", "file as target")
	pbFlagHeader *bool = flag.Bool("header", false, "output of the http-header")
)

func main() {

	flag.Parse()

	// returns the none-flag arguments only
	args:= flag.Args()

	var w io.Writer = os.Stdout

	if len(args) != 1 {
		fmt.Println("just min/max one url as input please")
		os.Exit(1)
	}
	
	for _, v := range args { fmt.Println(v) }

	url := args[0]

	resp, err := http.Get(url)
	if err != nil { 
		fmt.Println("shit happens url not found", url, err) 
		os.Exit(2)
	}
	// Response already loaded so dont forget to close
	defer resp.Body.Close()

	// creating path
	path := filepath.Dir(*pFlagOutput)
	if path != "" {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			fmt.Println("can't create path:", path, err)
			os.Exit(22)
		}
	}
	if *pFlagOutput != "" {
		pF, err := os.OpenFile(*pFlagOutput, os.O_CREATE | os.O_WRONLY, 0755)
		if err != nil {
			fmt.Println("can't creating file", *pFlagOutput, err )
			os.Exit(20)
		}
		defer pF.Close()

		w = pF
	}

	if *pbFlagHeader {
		//fmt.Printf("%#v", resp.Header)

		for k, v := range resp.Header {
			fmt.Println(k, v)
		}
	}

	// io.Copy(dsc io.Writer, src io.Reader)
	io.Copy(w, resp.Body)

}

func validateURL(s string) bool {

	_, err := url.ParseRequestURI(s)
	return err == nil	
	//validateURL := regexp.MustCompile("^http(s)?://[[:graph:]]+")	return validateURL.MatchString(s)
}