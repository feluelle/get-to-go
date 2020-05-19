package main

import (
	"exercises"
	"fmt"

	"os"

	"golang.org/x/tour/pic"
)

func main() {
	switch os.Args[1] {
	case "1":
		fmt.Println("Running Sqrt..")
		exercises.Sqrt(7)
	case "2":
		fmt.Println("Running Pic..")
		pic.Show(exercises.Pic)
	}
}
