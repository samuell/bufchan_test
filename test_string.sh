#!/bin/bash
go run test_string.go > output_string.txt
diff input.txt output_string.txt
