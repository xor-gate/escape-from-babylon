//go:build windows
// +build windows

package main

import (
	"C"
	"golang.org/x/sys/windows"
	"log"
	"os"
)

func systemGetWINEVersion() string {
	ntdll := windows.NewLazyDLL("ntdll.dll")
	wineGetVersionFunc := ntdll.NewProc("wine_get_version")

	err := wineGetVersionFunc.Find()
	if err != nil {
		return ""
	}

	r1, r2, r3 := wineGetVersionFunc.Call()

	log.Println("r1", r1)
	log.Println("r2", r2)
	log.Println("r3", r3)

	return ""
}

func systemIsUserRoot() bool {
	root := true

	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		root = false
	}

	return root
}
