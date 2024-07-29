//go:build !windows
// +build !windows

//
package main

func systemGetWINEVersion() string {
	return ""
}


func systemIsUserRoot() bool {
	return false
}
