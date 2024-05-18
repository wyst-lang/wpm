package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type PackageVersion struct {
	Author      string `json:"author"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type PackageIndex struct {
	Versions map[string]PackageVersion `json:"versions"`
}

func installPackage(packageName, packageVersion string) {
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

	if packageVersion == "" {
		packageVersion = getLatestVersion(pkgIndex)
	}

	pkgVersion, exists := pkgIndex.Versions[packageVersion]
	if !exists {
		fmt.Printf("Version %s not found for package %s\n", packageVersion, packageName)
		return
	}

	downloadAndSavePackage(pkgVersion.URL, packageName, packageVersion)
}

func getLatestVersion(pkgIndex PackageIndex) string {
	var latest string
	for version := range pkgIndex.Versions {
		if version > latest {
			latest = version
		}
	}
	return latest
}

func downloadAndSavePackage(url, packageName, packageVersion string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error downloading package: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: %s\n", resp.Status)
		return
	}

	libDir := "lib"
	if err := os.MkdirAll(libDir, 0755); err != nil {
		fmt.Printf("Error creating lib directory: %v\n", err)
		return
	}

	filename := filepath.Base(url)
	packageFile := filepath.Join(libDir, filename)
	out, err := os.Create(packageFile)
	if err != nil {
		fmt.Printf("Error creating package file: %v\n", err)
		return
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		fmt.Printf("Error saving package: %v\n", err)
	} else {
		fmt.Printf("Package %s version %s downloaded and saved successfully as %s\n", packageName, packageVersion, filename)
	}
}
