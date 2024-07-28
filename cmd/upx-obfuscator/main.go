package main

import (
	"os"
	"bytes"
	"log"
)

func bytesReplace(data, old, new []byte) []byte {
	foundIndex := bytes.Index(data, old)
	if foundIndex > -1 {
		// Found it!
		log.Println("Found identifier at offset", foundIndex)
	} else {
		return data
		log.Fatalln("Error file is not UPX packed")
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

	data = bytesReplace(data, []byte("UPX0"), []byte("GSP7"))
	data = bytesReplace(data, []byte("UPX1"), []byte("GSP1"))
	data = bytesReplace(data, []byte("UPX2"), []byte("GSP2"))

	_ = os.WriteFile(filename, data, 0666)

	log.Println("done")
}
