package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func rot13(input byte) byte {
	if 'A' <= input && input <= 'Z' {
		return 'A' + (input-'A'+13)%26
	} else if 'a' <= input && input <= 'z' {
		return 'a' + (input-'a'+13)%26
	}
	return input
}

func rot13Bytes(data []byte) []byte {
	result := make([]byte, len(data))
	for i, b := range data {
		result[i] = rot13(b)
	}
	return result
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <input file> <output file>\n", os.Args[0])
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read the input file
	inputData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	// Apply ROT13 transformation
	outputData := rot13Bytes(inputData)

	// Write the transformed data to the output file
	err = ioutil.WriteFile(outputFile, outputData, 0644)
	if err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}

	fmt.Printf("File %s has been converted using ROT13 and saved to %s\n", inputFile, outputFile)
}
