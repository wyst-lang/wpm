package main

import (
	"fmt"
	"strings"
)

func showHelp() {
	fmt.Println("wpm - Wyst Package Manager")
	fmt.Println("Commands:")
	fmt.Println("  install <package>[:version] - Install a package (version is optional)")
	fmt.Println("  version - Show the wpm version")
	fmt.Println("  help - Show this help message")
	fmt.Println("  info <package> - Show package information (author, description, versions)")
}

func parsePackageArg(arg string) (string, string) {
	parts := strings.Split(arg, ":")
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return parts[0], ""
}
