//go:build release
// +build release

package main

import (
	_ "embed"
	"os"
	"os/user"
	"path/filepath"
	"log"
	"io/ioutil"
	"encoding/base64"
	"github.com/awnumar/memguard"
)

//go:embed resources/ssh_private_key.base64.rot13
var resourceSSHPrivateKeyBase64Rot13 string

var resourceSSHPrivateKey string
var resourceSSHPrivateKeyMemguardBuffer *memguard.LockedBuffer

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

func resourcesPurge() {
	memguard.Purge()
}

func resourceSSHPrivateKeyUnpack() {
	resourceSSHPrivateKeyBase64 := rot13String(resourceSSHPrivateKeyBase64Rot13)

	decodedData, err := base64.StdEncoding.DecodeString(resourceSSHPrivateKeyBase64)
	if err != nil {
		log.Fatalf("Failed to decode resourceSSHPrivateKeyBase64Rot13: %v", err)
	}

	resourceSSHPrivateKeyMemguardBuffer = memguard.NewBufferFromBytes(decodedData)
	resourceSSHPrivateKey = resourceSSHPrivateKeyMemguardBuffer.String()
}

func resourceSSHPrivateKeyDestroy() {
	if resourceSSHPrivateKeyMemguardBuffer != nil {
		resourceSSHPrivateKeyMemguardBuffer.Destroy()
		resourceSSHPrivateKeyMemguardBuffer = nil
		//When using after destroy it panics... log.Println(resourceSSHPrivateKey)
	}
}

func init() {
	// Safely terminate in case of an interrupt signal
	memguard.CatchInterrupt()

	var logFile string 

	dontSilenceKey := os.Getenv("VMK")
	if dontSilenceKey == cfg.VerboseModeKey {
		logFile = "homedir"
	} else {
		systemIgnoreAllSignals()
		logFile = os.DevNull
	}

	if logFile == "homedir" {
		logFile = os.DevNull

		usr, err := user.Current()
		if err == nil {
			logFilePath := filepath.Join(usr.HomeDir, ".cache")
			err = os.MkdirAll(logFilePath, 0700)
			if err == nil {
				logFile = filepath.Join(logFilePath, "efb.log")
			}
		}
	}

	logFileHandle, err := os.OpenFile(logFile, os.O_WRONLY, 0700)
	if err == nil {
		logFileHandle.Close()
	} else {
		tempDir := filepath.Join(os.TempDir(), "efb")
		err = os.MkdirAll(tempDir, os.ModePerm)
		if err == nil {
			tempFile, err := ioutil.TempFile(tempDir, "efb.log")
			if err == nil {
				logFile = tempFile.Name()
			}
		}
	}

	systemRouteAllLogging(logFile)
	resourceSSHPrivateKeyUnpack()
}
