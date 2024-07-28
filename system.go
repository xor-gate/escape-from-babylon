package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
)

// Route all logging
func systemRouteAllLogging(logfile string) {
	nullFile, err := os.OpenFile(logfile, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error opening /dev/null:", err)
		return
	}

	// Redirect stdout and stderr to /dev/null
	os.Stdout = nullFile
	os.Stderr = nullFile

	// Redirect log facility to /dev/null
	log.SetOutput(nullFile)
}

func systemGetAppDataPath() string {
	return filepath.Join(os.Getenv("USERPROFILE"), "AppData")
}

// systemCheckDirExists checks if the directory at the given path exists.
func systemIsDirExisting(path string) bool {
	// Get file info
	info, err := os.Stat(path)
	if err != nil {
		// If the error is due to the file not existing, return false
		if os.IsNotExist(err) {
			return false
		}
		// For any other errors, you may want to handle them as needed
		return false
	}

	// Check if the info corresponds to a directory
	return info.IsDir()
}

func systemGetFilesInDirectory(path string) ([]string, bool) {
	var filesInDirectory []string

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, false
	}

	for _, file := range files {
		filesInDirectory = append(filesInDirectory, file.Name())
	}

	return filesInDirectory, true
}

func systemIsFileExisting(path string) bool {
	// Get file info
	info, err := os.Stat(path)
	if err != nil {
		// If the error is due to the file not existing, return false
		if os.IsNotExist(err) {
			return false
		}
		// For any other errors, you may want to handle them as needed
		return false
	}

	// Check if the info corresponds to a regular file
	return !info.IsDir()
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

func systemGetWellKnownExistingPaths() []string {
	var existingPaths []string

	appDataPath := systemGetAppDataPath()
	if ok := systemIsDirExisting(appDataPath); !ok {
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

func systemIgnoreAllSignals() {
	// Create a channel to receive OS signals.
	sigs := make(chan os.Signal, 1)

	// Notify the signal channel for all signals (you can add more if needed)
	signal.Notify(sigs)

	// This goroutine receives signals but does nothing with them.
	go func() {
		for sig := range sigs {
			// Signal received but ignored
			_ = sig
			log.Println("Received OS signal", sig)
		}
	}()
}

func systemOSDetect() {
	wineOSFiles := systemGetWellKnownWINEOSFiles()
	if len(wineOSFiles) != 0 {
		log.Println("WINE detected")
		for _, file := range wineOSFiles {
			log.Println("\t", file)
		}
	}
}
