package main

import (
	"fmt"
	"net/http"
	"strings"
)

const URL = "http://localhost:3000"

func showHelp() {
	fmt.Println("wpm - Wyst Package Manager")
	fmt.Println("Commands:")
	fmt.Println("  install <package>[:version] - Install a package (version is optional)")
	fmt.Println("  create <package> - Create a package")
	fmt.Println("  edit <package> - Edit a package")
	fmt.Println("  delete <package> - Delete a package")
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

func createPackage(packageName string) {
	var psw string
	var pswV string
	var repo string
	var version string
	fmt.Printf("Enter the password for %s: ", packageName)
	ReadPassword(&psw)
	fmt.Printf("Verify the password for %s: ", packageName)
	ReadPassword(&pswV)
	if psw != pswV {
		fmt.Printf("Passwords do not match\n\r")
		return
	}
	fmt.Printf("Enter the repo for %s: ", packageName)
	fmt.Scanln(&repo)
	fmt.Printf("Enter the latest version for %s: ", packageName)
	fmt.Scanln(&version)
	fmt.Println()
	jsonBody := []byte(fmt.Sprintf(`{"name": "%v", "psw": "%v", "repo": "%v", "latest": "%v"}`, packageName, psw, repo, version))
	req, err := sendRequest(http.MethodPost, URL, jsonBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = getMessage(req)
	if err != nil {
		fmt.Printf("Error: %v\n\r", err)
		return
	} else {
		fmt.Printf("Done\n\r")
	}
}

func editPackage(packageName string) {
	pkgidx, err := getPackage(packageName)
	if err != nil {
		fmt.Println(err)
		return
	}
	var psw string
	var repo string = pkgidx.Repo
	var version string = pkgidx.Latest
	var newPsw string
	var newPackageName string = packageName

	fmt.Printf("\rEnter the password for %s: ", packageName)
	ReadPassword(&psw)
	newPsw = psw

	for true {
		clear()
		fmt.Printf("Choose what you want to modify\n\r")
		fmt.Printf("  1 -> package name  : %s\n\r", newPackageName)
		fmt.Printf("  2 -> password      : %s\n\r", newPsw)
		fmt.Printf("  3 -> repo          : %s\n\r", repo)
		fmt.Printf("  4 -> latest version: %s\n\r", version)
		fmt.Printf("Type in 'quit'/'q'/'exit' to discard your changes\n\r")
		fmt.Printf("Type in anything else to confirm your changes\n\r")
		var option string
		fmt.Print("\n> ")
		fmt.Scanln(&option)
		switch option {
		case "1":
			fmt.Printf("Enter the new name for for %s: ", packageName)
			fmt.Scanln(&newPackageName)
		case "2":
			fmt.Printf("Enter the new password for %s: ", packageName)
			ReadPassword(&newPsw)
		case "3":
			fmt.Printf("Enter the repo for %s: ", packageName)
			fmt.Scanln(&repo)
		case "4":
			fmt.Printf("Enter the latest version for %s: ", packageName)
			fmt.Scanln(&version)
		case "q":
			return
		case "quit":
			return
		case "exit":
			return
		default:
			jsonBody := []byte(fmt.Sprintf(`{ "new": {"name": "%v", "psw": "%v", "latest": "%v", "repo": "%v"}, "name": "%v", "psw": "%v"}`, newPackageName, newPsw, version, repo, packageName, psw))
			req, err := sendRequest(http.MethodPut, URL, jsonBody)
			if err != nil {
				fmt.Printf("Error while editing %s: %v", packageName, err)
				return
			}
			err = getMessage(req)
			if err != nil {
				fmt.Printf("Error: %v\n\r", err)
				return
			} else {
				fmt.Printf("Done\n\r")
			}
			return
		}

	}
}

func deletePackage(packageName string) {
	fmt.Printf("Enter the password for %s: ", packageName)
	var psw string
	ReadPassword(&psw)
	jsonBody := []byte(fmt.Sprintf(`{"name": "%s", "psw": "%v"}`, packageName, psw))
	req, err := sendRequest(http.MethodDelete, URL, jsonBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = getMessage(req)
	if err != nil {
		fmt.Printf("Error: %v\n\r", err)
		return
	} else {
		fmt.Printf("Done\n\r")
	}
}
