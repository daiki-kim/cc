package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var flagByte bool
	flag.BoolVar(&flagByte, "c", false, "print byte of file contents")
	flag.Parse()

	fName := flag.Arg(0)
	if fName == "" {
		log.Fatalln("need file name")
	}
	file, err := os.Open(fName)
	if err != nil {
		log.Fatalln("failed to open file:", err)
	}

	if flagByte {
		fBytes, err := io.ReadAll(file)
		if err != nil {
			log.Fatalln("failed to read file:", err)
		}
		fmt.Println(len(fBytes))
	}
}
