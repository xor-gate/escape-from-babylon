//go:build darwin
// +build darwin

package main

func systemGetWINEVersion() string {
	return ""
}

func systemGetUname() {
}

func systemIsUserRoot() bool {
	return false
}
