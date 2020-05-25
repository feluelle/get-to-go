package main

import (
	"errors"
	"strconv"
)

// Bin2Dec Converts binary to decimal
// See https://github.com/florinpop17/app-ideas/blob/master/Projects/1-Beginner/Bin2Dec-App.md
func Bin2Dec(binaryNumber string, length int) (int64, error) {
	if len(binaryNumber) > length {
		return -1, errors.New("Binary number exceeds allowed length")
	}
	decimalNumber, _ := strconv.ParseInt(binaryNumber, 2, 64)
	return decimalNumber, nil
}
