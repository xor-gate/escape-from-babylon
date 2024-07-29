//go:build dll
// +build dll

package main

import (
	"C"
)

//export runMain
func runMain() {
	main()
}
