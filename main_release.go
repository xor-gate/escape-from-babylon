//go:build release
// +build release
package main

import (
	_ "embed"
	"os"
	"fmt"
	"log"
)

//go:embed resources/ssh_private_key
var resourceSSHPrivateKey string

func init() {
	// Open /dev/null for writing
	nullFile, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error opening /dev/null:", err)
		return
	}

	// Redirect stdout and stderr to /dev/null
	os.Stdout = nullFile
	os.Stderr = nullFile

	// Redirect log facility to /dev/null
	log.SetOutput(nullFile)
}
