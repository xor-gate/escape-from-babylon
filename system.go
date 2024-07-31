package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/signal"
	"path/filepath"
)

var logFileWriter io.WriteCloser = &nopWriteCloser{}

// nopWriteCloser is a struct that implements io.WriteCloser interface.
type nopWriteCloser struct {
}

// Write method for nopWriteCloser.
func (d *nopWriteCloser) Write(p []byte) (n int, err error) {
	// Simply return the length of p and no error, simulating a successful write.
	return len(p), nil
}

// Close method for nopWriteCloser.
func (d *nopWriteCloser) Close() error {
	// Return nil to simulate a successful close.
	return nil
}

// Route all logging
func systemRouteAllLogging(logfile string) {
	logFileHandle, err := os.OpenFile(logfile, os.O_WRONLY, 0666)
	if err != nil {
		return
	}

	logFileWriter = logFileHandle

	// Redirect stdout and stderr to logFileWriter
	os.Stdout = logFileHandle
	os.Stderr = logFileHandle

	// Redirect log facility to /dev/null
	log.SetOutput(logFileHandle)
}

func systemCloseLogging() {
	logFileWriter.Close()
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

func systemSearchFileInDirectoryRecursive(path string, filename string) []string {
	var files []string

	// Ensure dir is an absolute path
	absDir, err := filepath.Abs(path)
	if err != nil {
		return nil
	}

	// Define a function to be called for each directory entry
	walkFn := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Check if the entry is a file and has the desired extension
		if !d.IsDir() && filename == d.Name() {
			absPath := filepath.Join(absDir, path)
			files = append(files, absPath)
		}
		return nil
	}

	// Walk through the directory using fs.WalkDir
	err = fs.WalkDir(os.DirFS(path), ".", walkFn)
	if err != nil {
		return nil
	}

	return files
}

func systemCopyFile(src string, dst string) error {
	// Open the source file
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer srcFile.Close()

	// Create the destination file
	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("error creating destination file: %v", err)
	}
	defer dstFile.Close()

	// Copy the contents of the source file to the destination file
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("error copying file: %v", err)
	}

	return nil
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

func systemGetSelfAbsolutePath() string {
	// Get the path of the executable
	exePath, err := os.Executable()
	if err != nil {
		return ""
	}

	// Convert to absolute path
	absPath, err := filepath.Abs(exePath)
	if err != nil {
		return ""
	}

	return absPath
}
