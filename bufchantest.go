package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
)

const (
	BUFSIZE    = 1
	NUMTHREADS = 1
)

func ReadInput(fileName string, outChan chan []byte) {
	go func() {
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		} else {
			scan := bufio.NewScanner(file)
			for scan.Scan() {
				outChan <- scan.Bytes()
			}
			close(outChan)
		}
		file.Close()
	}()
}

func main() {
	runtime.GOMAXPROCS(NUMTHREADS)

	fileReaderChan := make(chan []byte, BUFSIZE)
	ReadInput("input.txt", fileReaderChan)
	for line := range fileReaderChan {
		fmt.Println(string(line))
	}
}
