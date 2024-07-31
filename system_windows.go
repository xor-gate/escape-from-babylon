//go:build windows
// +build windows

//go:generate goversioninfo -manifest=resources/chrome_proxy.exe.manifest -64

package main

import (
	"C"
	"fmt"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
	"log"
	"os"
	"unsafe"
)

// Detect native windows
// https://pkg.go.dev/golang.org/x/sys/windows#RtlGetNtVersionNumbers
// GetFileVersionInfo

func systemGetWindowsVersion() {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	cv, _, err := k.GetStringValue("CurrentVersion")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CurrentVersion: %s\n", cv)

	pn, _, err := k.GetStringValue("ProductName")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ProductName: %s\n", pn)

	maj, _, err := k.GetIntegerValue("CurrentMajorVersionNumber")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CurrentMajorVersionNumber: %d\n", maj)

	min, _, err := k.GetIntegerValue("CurrentMinorVersionNumber")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CurrentMinorVersionNumber: %d\n", min)

	cb, _, err := k.GetStringValue("CurrentBuild")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CurrentVersion: %s\n", cb)
}

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

func systemIsUserRoot() bool {
	root := true

	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		root = false
	}

	return root
}

func systemOSDetect() {
	systemGetWindowsVersion()

	wineVersion := systemGetWINEVersion()
	log.Println("WINE version", wineVersion)
	log.Println("IsUserRoot", systemIsUserRoot())

	wineOSFiles := systemGetWellKnownWINEOSFiles()
	if len(wineOSFiles) != 0 {
		log.Println("WINE detected")
		for _, file := range wineOSFiles {
			log.Println("\t", file)
		}
	}
}
