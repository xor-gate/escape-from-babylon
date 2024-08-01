package main

import (
	"os"
	"bytes"
	"log"
	"fmt"
)

func bytesReplace(data, old, new []byte) []byte {
	foundIndex := bytes.Index(data, old)
	if foundIndex > -1 {
		// Found it!
		log.Println("Found identifier at offset", foundIndex)
	} else {
		return data
	}

	return bytes.Replace(data, old, new, 1)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Specify exe file to obfuscate")
	}

	filename := os.Args[1]

	log.Println("Obfuscating UPX compressed executable file")
	log.Println("\t", filename)	

	data, _ := os.ReadFile(filename)

	for i := range(10) {
		upxIdentifier := fmt.Sprintf("UPX%d", i)
		efbIdentifier := fmt.Sprintf("EFB%d", i)
		data = bytesReplace(data, []byte(upxIdentifier), []byte(efbIdentifier))
	}

	_ = os.WriteFile(filename, data, 0666)

	log.Println("done")
}
