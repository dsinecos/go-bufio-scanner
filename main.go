package main

import (
	"fmt"
	"strings"
)

func main() {

	inputString := "John Jane Alex"
	stringReader := strings.NewReader(inputString)

	fmt.Println(inputString)
	fmt.Println(stringReader.Len())
	fmt.Println(stringReader.Size())

	byteReader := make([]byte, 10)

	stringReader.Read(byteReader)
	fmt.Println(string(byteReader))
	fmt.Println(stringReader.Len())
	fmt.Println(stringReader.Size())

	readByte, err := stringReader.ReadByte()

	if err != nil {
		panic(err)
	}

	fmt.Println(string(readByte))

	splitString := strings.Split(inputString, " ")
	for _, element := range splitString {
		fmt.Println(element)
	}
}
