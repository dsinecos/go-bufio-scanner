package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func main() {

	// Read a fixed number of bytes each time from the file
	// The number of bytes read can be defined as a constant

	const MaxLength = 10
	inputString := "abcdefghijklmnopqrstuvwxyz"
	stringScanner := bufio.NewScanner(strings.NewReader(inputString))

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// fmt.Printf("%t\t%d\t%s\n", atEOF, len(data), data)

		if len(data) <= MaxLength && len(data) != 0 {
			return len(data), data, nil
		}

		return 0, nil, nil
	}

	stringScanner.Split(split)

	buf := make([]byte, 10)
	stringScanner.Buffer(buf, bufio.MaxScanTokenSize)

	for stringScanner.Scan() {
		color.Cyan("Inside scan")
		fmt.Printf("%s\n", stringScanner.Text())
	}

	errScanning := stringScanner.Err()

	if errScanning != nil {
		color.Yellow(errScanning.Error())
	}

}
