package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Readln reads a line from stdin
func Readln() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// DownloadFile ...
func DownloadFile(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}
