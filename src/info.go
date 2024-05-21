package main

import "fmt"

func showPackageInfo(packageName string) {
	progress := ProgressBar{total: 4, length: 20}
	progress.change(3, "", "Fetching")
	pkgidx, err := getPackage(packageName)
	if err != nil {
		fmt.Printf("\nError getting repo URL: %v\n", err)
		return
	}
	progress.clean()

	// indexURL, err := findIndexURL(repoURL)
	// if err != nil {
	// 	fmt.Printf("Error finding index URL: %v\n", err)
	// 	return
	// }

	// pkgIndex, err := fetchPackageIndex(indexURL)
	// if err != nil {
	// 	fmt.Printf("Error fetching package index: %v\n", err)
	// 	return
	// }

	fmt.Printf("Package: %s\n", packageName)
	fmt.Printf("Repo: %s\n", pkgidx.Repo)
	fmt.Printf("Version: %s\n", pkgidx.Latest)
	// for version, info := range pkgIndex.Versions {
	// fmt.Printf("Version: %s\n", version)
	// fmt.Printf("  Author: %s\n", info.Author)
	// fmt.Printf("  Description: %s\n", info.Description)
	// fmt.Println()
	// }
}
