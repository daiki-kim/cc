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
	flag.BoolVar(&flagByte, "b", false, "print byte count of file contents")
	flag.BoolVar(&flagLine, "l", false, "print line count of file contents")
	flag.BoolVar(&flagWord, "w", false, "print word count of file contents")
	flag.BoolVar(&flagChar, "c", false, "print character count of file contents")
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
		file.Seek(0, 0)
		fmt.Println(len(fBytes))
	}

	if flagLine {
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
		fmt.Println(lCnt)
	}

	if flagWord {
		wCnt := 0
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			wCnt++
		}
		file.Seek(0, 0)
		fmt.Println(wCnt)
	}

	if flagChar {
		cCnt := 0
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanRunes)
		for scanner.Scan() {
			cCnt++
		}
		file.Seek(0, 0)
		fmt.Println(cCnt)
	}
}
