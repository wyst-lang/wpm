package main

import (
	"fmt"
)

func showPackageInfo(packageName string) {
	repoURL, err := getRepoURL(packageName)
	if err != nil {
		fmt.Printf("Error getting repo URL: %v\n", err)
		return
	}

	indexURL, err := findIndexURL(repoURL)
	if err != nil {
		fmt.Printf("Error finding index URL: %v\n", err)
		return
	}

	pkgIndex, err := fetchPackageIndex(indexURL)
	if err != nil {
		fmt.Printf("Error fetching package index: %v\n", err)
		return
	}

	fmt.Printf("Package: %s\n", packageName)
	for version, info := range pkgIndex.Versions {
		fmt.Printf("Version: %s\n", version)
		fmt.Printf("  Author: %s\n", info.Author)
		fmt.Printf("  Description: %s\n", info.Description)
		fmt.Println()
	}
}
