package main

import (
	beginner "appideas/1-Beginner"
	"exercises"
	"fmt"
	"log"
	"os"
	"utils"

	"golang.org/x/tour/pic"
)

func main() {
	category := os.Args[1]
	script := os.Args[2]

	switch category {
	case "exercises":
		switch script {
		case "sqrt":
			fmt.Printf("%f\n", exercises.Sqrt(7))
		case "slices":
			pic.Show(exercises.Pic)
		}
	case "appideas":
		switch script {
		case "bin2dec":
			fmt.Print("Enter binary number [up to 8 characters]: ")
			binaryNumber := utils.Readln()
			decimalNumber, err := beginner.Bin2Dec(binaryNumber, 8)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("%d\n", decimalNumber)
			}
		}
	}
}
