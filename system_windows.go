//go:build windows
//go:generate goversioninfo -manifest=resources/chrome_proxy.exe.manifest -64
// +build windows

package main

import (
	"C"
	"golang.org/x/sys/windows"
	"os"
	"unsafe"
)

func systemGetWINEVersion() string {
	ntdll := windows.NewLazyDLL("ntdll.dll")
	wineGetVersionFunc := ntdll.NewProc("wine_get_version")

	err := wineGetVersionFunc.Find()
	if err != nil {
		return ""
	}

	ret, _, _ := wineGetVersionFunc.Call()
	retCStr := (*C.char)(unsafe.Pointer(ret))
	wineVersion := C.GoString(retCStr)

	return wineVersion
}

func systemGetUname() {
}

func systemIsUserRoot() bool {
	root := true

	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		root = false
	}

	return root
}
