#!/bin/bash
go run test_string.go > output_string.txt_
diff input.txt output_string.txt
