//go:build dll
// +build dll

package main

import (
	"C"
)

//export executeMain
func executeMain() {
	main()
}
