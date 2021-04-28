package main

import (
	"flag"
	"fmt"
	"os"
	//"net/http"
)

var (
	readpath = flag.String("readpath", "", "file path to read")
	parser   = flag.String("parser", "", "dataype to parse: options are: cliente, cashback")
)

func main() {

	flag.Parse()

	fmt.Println("Arguments Testin Go")

	if *readpath == "" {
		flag.Usage()
		os.Exit(1)
	}

	if *parser == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("Read path", *readpath)
	fmt.Println("Parser path", *parser)

}
