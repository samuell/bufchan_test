#!/bin/bash
go run test_bytearray.go > output_bytearray.txt
diff input.txt output_bytearray.txt
