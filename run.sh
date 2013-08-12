#!/bin/bash
go run bufchantest.go > output.txt
diff input.txt output.txt
