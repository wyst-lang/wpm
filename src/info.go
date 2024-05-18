package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func showPackageInfo(packageName string) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/wyst-lang/index/master/%s/index.json", packageName)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching package info: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: %s\n", resp.Status)
		return
	}

	var pkgIndex PackageIndex
	if err := json.NewDecoder(resp.Body).Decode(&pkgIndex); err != nil {
		fmt.Printf("Error decoding package info: %v\n", err)
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
