package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

const (
	BUFSIZE    = 1
	NUMTHREADS = 1
)

func ReadInput(fileName string, ch chan string) {
	go func() {
		file, err := os.Open(fileName)
		if err == nil {
			scan := bufio.NewScanner(file)
			for scan.Scan() {
				ch <- scan.Text()
			}
			close(ch)
			file.Close()
		}
	}()
}

func main() {
	runtime.GOMAXPROCS(NUMTHREADS)
	ch := make(chan string, BUFSIZE)
	ReadInput("input.txt", ch)
	for line := range ch {
		fmt.Println(string(line))
	}
}
