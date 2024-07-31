//go:build linux
// +build linux

package main

import (
	"log"
	"strings"
	"syscall"
)

func systemGetWINEVersion() string {
	return ""
}

func systemGetUname() {
	var uts syscall.Utsname

	err := syscall.Uname(&uts)
	if err != nil {
		log.Println("Error getting system information:", err)
		return
	}

	// Convert the byte arrays to strings
	sysname := int8SliceToString(uts.Sysname[:])
	release := int8SliceToString(uts.Release[:])
	version := int8SliceToString(uts.Version[:])

	// Check for FreeBSD Linux emulation specific indicators
	log.Println("syscall.Uname:", "(sysname)", sysname, "(release)", release, "(version)", version)
	if strings.Contains(sysname, "Linux") && strings.Contains(version, "FreeBSD") {
		log.Println("Running under FreeBSD linuxemu")
	}
}

// int8SliceToString converts a slice of int8 to a string.
func int8SliceToString(int8Slice []int8) string {
	// Create a byte slice with the same length as the int8 slice
	byteSlice := make([]byte, len(int8Slice))
	for i, v := range int8Slice {
		byteSlice[i] = byte(v)
	}
	return string(byteSlice)
}

func systemIsUserRoot() bool {
	return false
}
