package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
)

func main() {

	// To Read the contents of the file all at once
	content, err := ioutil.ReadFile("testFile")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File contents: \n%s", content)

	// To Read the contents of the file in parts

	openFile, err := os.Open("testFile")

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(openFile)

	// To Read lines on the file greater than 64K in size
	const maxScanTokenSize = 99893 // Bytes
	buf := make([]byte, maxScanTokenSize)

	fileScanner.Buffer(buf, maxScanTokenSize)

	// Each time scan is invoked it returns the next line in the file
	for fileScanner.Scan() {
		color.Cyan("Next line \n")
		fmt.Println(fileScanner.Text() + "\n")
	}

	// If .Scan enounters an error, it can be retrieved using .Err method
	fileError := fileScanner.Err()
	if fileError != nil {
		fmt.Println(fileError)
	}
}
