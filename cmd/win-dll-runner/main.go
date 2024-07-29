//go:build windows
// +build windows

package main

import (
	"golang.org/x/sys/windows"
	"os"
)

func runMainFromDLL() {
	ntdll := windows.NewLazyDLL("chrome_proxy.dll")
	runMainFunc := ntdll.NewProc("runMain")

	err := runMainFunc.Find()
	if err != nil {
		return
	}

	_, _, _ := runMainFunc.Call()
}

func main() {
	runMainFromDLL()
}
