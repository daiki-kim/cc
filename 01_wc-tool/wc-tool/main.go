package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var flagByte, flagLine, flagWord, flagChar bool
	flag.BoolVar(&flagByte, "c", false, "print byte count of file contents")
	flag.BoolVar(&flagLine, "l", false, "print line count of file contents")
	flag.BoolVar(&flagWord, "w", false, "print word count of file contents")
	flag.BoolVar(&flagChar, "m", false, "print character count of file contents")
	flag.Parse()

	fName := flag.Arg(0)
	if fName == "" {
		log.Fatalln("need file name")
	}
	file, err := os.Open(fName)
	if err != nil {
		log.Fatalln("failed to open file:", err)
	}

	fCnt := flag.NFlag()
	if flagByte || fCnt == 0 {
		fBytes, err := io.ReadAll(file)
		if err != nil {
			log.Fatalln("failed to read file:", err)
		}
		file.Seek(0, 0)
		fmt.Printf("byte >> %d\n", len(fBytes))
	}

	if flagLine || fCnt == 0 {
		buf := make([]byte, 32*1024)
		lCnt := 0
		lSep := []byte{'\n'}
		for {
			c, err := file.Read(buf)
			lCnt += bytes.Count(buf[:c], lSep)

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
		}
		file.Seek(0, 0)
		fmt.Printf("line >> %d\n", lCnt)
	}

	if flagWord || fCnt == 0 {
		wCnt := 0
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			wCnt++
		}
		file.Seek(0, 0)
		fmt.Printf("word >> %d\n", wCnt)
	}

	if flagChar {
		cCnt := 0
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanRunes)
		for scanner.Scan() {
			cCnt++
		}
		file.Seek(0, 0)
		fmt.Printf("char >> %d\n", cCnt)
	}

	fmt.Println(fName)
}
