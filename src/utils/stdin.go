package utils

import (
	"bufio"
	"os"
)

// Readln reads a line from stdin
func Readln() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
