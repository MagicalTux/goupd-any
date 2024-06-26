package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/KarpelesLab/goupd"
)

func main() {
	exe, err := os.Executable()
	if err != nil {
		exe = os.Args[0]
	}
	exe = strings.TrimSuffix(strings.ToLower(filepath.Base(exe)), ".exe")

	if exe == "goupd-any" {
		log.Printf("invalid package name\n")
		os.Exit(1)
		return
	}

	log.Printf("using project name %s", exe)

	goupd.PROJECT_NAME = exe
	goupd.RunAutoUpdateCheck()

	log.Printf("failed to get version\n")
	os.Exit(1)
	return
}
