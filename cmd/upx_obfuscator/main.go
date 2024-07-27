package main

import (
	"os"
	"bytes"
	"log"
)

var originalIdentifier = []byte("UPX0")	 
var obfuscatedIdentifier = []byte("GSP7")	 

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Specify exe file to obfuscate")
	}

	filename := os.Args[1]

	log.Println("Obfuscating UPX compressed executable file")
	log.Println("\t", filename)	

	data, _ := os.ReadFile(filename)

	foundIndex := bytes.Index(data, originalIdentifier)
	if foundIndex > -1 {
		// Found it!
		log.Println("Found UPX identifier at offset", foundIndex)
	} else {
		log.Fatalln("Error file is not UPX packed")
	}


	obfuscatedData := bytes.Replace(data, originalIdentifier, obfuscatedIdentifier, 1)
	_ = os.WriteFile(filename, obfuscatedData, 0666)

	log.Println("done")
}
