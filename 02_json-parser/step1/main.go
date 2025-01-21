package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fName := os.Args[1]
	f, err := os.Open(fName)
	if err != nil {
		log.Fatalln("failed to open file:", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	firstCharFound := false
	openBraceFound := false
	closeBraceFound := false
	for scanner.Scan() {
		char := scanner.Text()

		if !firstCharFound && (char == " " || char == "\t" || char == "\r" || char == "\n") {
			continue
		}

		if !firstCharFound {
			log.Println("first char found")
			firstCharFound = true
			if char != "{" {
				log.Fatalln("json must start with '{'")
			}
			openBraceFound = true
			continue
		}

		if char != " " && char != "\t" && char != "\r" && char != "\n" {
			if char == "{" {
				if openBraceFound {
					fmt.Println("JSON is invalid")
					os.Exit(1)
				}
				openBraceFound = true
			} else if char == "}" {
				if closeBraceFound {
					fmt.Println("JSON is invalid")
					os.Exit(1)
				}
				if openBraceFound {
					closeBraceFound = true
				}
			} else if openBraceFound && closeBraceFound {
				openBraceFound = false
				closeBraceFound = false
			}
		}
	}

	if openBraceFound && closeBraceFound {
		fmt.Println("JSON is valid")
		os.Exit(0)
	}

	fmt.Println("JSON is invalid")
	os.Exit(1)
}

// func main() {
// 	fName := os.Args[1]
// 	f, err := os.Open(fName)
// 	if err != nil {
// 		log.Fatalln("failed to open file:", err)
// 	}
// 	defer f.Close()

// 	fData, err := io.ReadAll(f)
// 	if err != nil {
// 		log.Fatalln("failed to read file:", err)
// 	}

// 	if string(fData[:1]) == "{" && string(fData[len(fData)-1:]) == "}" {
// 		fmt.Println("JSON is valid")
// 		os.Exit(0)
// 	}
// 	fmt.Println("JSON is invalid")
// 	os.Exit(1)
// }
