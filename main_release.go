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
	systemSilenceAllLogging()
}
