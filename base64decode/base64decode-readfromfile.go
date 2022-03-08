package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	encodedfile, err := ioutil.ReadFile("encodedfile")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	encodedstring := string(encodedfile)
	fmt.Println(encodedstring)

	fmt.Println(encodedstring)

	dec, err := base64.StdEncoding.DecodeString(encodedstring)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("decodedfile")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}
}
