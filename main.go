package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/magicaltux/goupd"
)

func main() {
	exe, err := os.Executable()
	if err != nil {
		exe = os.Args[0]
	}
	exe = strings.TrimSuffix(strings.ToLower(path.Base(exe)), ".exe")

	if exe == "goupd-any" {
		fmt.Printf("invalid package name\n")
		os.Exit(1)
		return
	}

	goupd.PROJECT_NAME = exe
	goupd.RunAutoUpdateCheck()

	fmt.Printf("failed to get version\n")
	os.Exit(1)
	return
}
