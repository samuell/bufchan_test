#!/bin/bash
go run test_bytearray_wochan.go > output_bytearray_wochan.txt
diff input.txt output_bytearray_wochan.txt
