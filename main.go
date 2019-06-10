package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

func main() {

	// Read a fixed number of bytes each time from the file
	// The number of bytes read can be defined as a constant

	const MaxLengthPerToken = 10
	openFile, err := os.Open("testFile")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(openFile)

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if !atEOF {
			fmt.Println("This is the End Of File")
			fmt.Println(string(data))
		}
		if len(data) <= MaxLengthPerToken && len(data) != 0 {
			return len(data), data, nil
		}
		return 0, nil, nil
	}

	fileScanner.Split(split)

	buf := make([]byte, MaxLengthPerToken)
	fileScanner.Buffer(buf, MaxLengthPerToken)

	for fileScanner.Scan() {
		color.Cyan("Scanning next token")
		fmt.Printf("%s\n", fileScanner.Text())
	}

	errScanning := fileScanner.Err()

	if errScanning != nil {
		color.Yellow(errScanning.Error())
	}

}
