package main

import (
	"fmt"
	"os"
)

const wpmVersion = "1.0.0"

func main() {
	fmt.Println()
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	command := os.Args[1]
	switch command {
	case "install":
		if len(os.Args) < 3 {
			fmt.Println("Usage: wpm install <package>[:version]")
			return
		}
		packageName, packageVersion := parsePackageArg(os.Args[2])
		installPackage(packageName, packageVersion)
	case "version":
		fmt.Println("wpm version", wpmVersion)
	case "help":
		showHelp()
	case "info":
		if len(os.Args) < 3 {
			fmt.Println("Usage: wpm info <package>")
			return
		}
		showPackageInfo(os.Args[2])
	default:
		fmt.Println("Unknown command:", command)
		showHelp()
	}
}
