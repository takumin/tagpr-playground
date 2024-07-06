package main

import (
	"fmt"
	"runtime"

	"github.com/takumin/tagpr-playground/internal/version"
)

func main() {
	showVersion()
}

func showVersion() {
	fmt.Println("Runtime:", runtime.Version())
	fmt.Println("Version:", version.Version())
	fmt.Println("Revision:", version.Revision())
	fmt.Println("Prerelease:", version.Prerelease())
}

func testFunc() {
	fmt.Println("test implemented")
}
