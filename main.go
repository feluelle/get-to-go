package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	script := os.Args[1]
	switch script {
	case "bin2dec":
		fmt.Print("Enter binary number [up to 8 characters]: ")
		binaryNumber := Readln()
		decimalNumber, err := Bin2Dec(binaryNumber, 8)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Decimal number: %d\n", decimalNumber)
		}
	case "emojitranslator":
		fmt.Print("Enter a text to translate to emojis: ")
		text := Readln()
		emojis, codes := EmojiTranslator(text)
		fmt.Printf("Codes: %v\n", codes)
		fmt.Printf("Emojis: %s\n", emojis)
	}
}
