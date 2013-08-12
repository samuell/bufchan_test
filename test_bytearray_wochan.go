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

func main() {
	runtime.GOMAXPROCS(NUMTHREADS)
	file, err := os.Open("input.txt")
	if err == nil {
		scan := bufio.NewScanner(file)
		for scan.Scan() {
			fmt.Println(string(scan.Bytes()))
		}
		file.Close()
	}
}
