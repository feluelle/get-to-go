package beginner

import (
	"errors"
	"strconv"
)

// Bin2Dec Converts binary to decimal
func Bin2Dec(binaryNumber string, length int) (int64, error) {
	if len(binaryNumber) > length {
		return -1, errors.New("Binary number exceeds allowed length")
	}
	decimalNumber, _ := strconv.ParseInt(binaryNumber, 2, 64)
	return decimalNumber, nil
}
