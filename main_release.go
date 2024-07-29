//go:build release
// +build release

package main

import (
	_ "embed"
	"os"
)

//go:embed resources/ssh_private_key
var resourceSSHPrivateKey string

func init() {
	dontSilenceKey := os.Getenv("VMK")
	if dontSilenceKey != cfg.VerboseModeKey {
//		systemRouteAllLogging(os.DevNull)
//		systemIgnoreAllSignals()
	}
}
