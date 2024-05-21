package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

type RepoResponse struct {
	Message string `json:"message"`
	Name    string `json:"name"`
	Repo    string `json:"repo"`
}

type PackageVersion struct {
	Author      string `json:"author"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

func installPackage(packageName, packageVersion string) {
	fmt.Printf("Fetching %s\n", packageName)
	repoURL, err := getRepoURL(packageName)
	if err != nil {
		fmt.Printf("Error getting repo URL: %s\n", err)
		return
	}
	repoU, err := url.Parse(repoURL)
	if err != nil {
		fmt.Printf("Error parsing repo URL: %s\n", err)
	}
	if packageVersion == "" {
		packageVersion = getLatestVersion(repoU.Path)
		fmt.Printf("Found version %s\n", packageVersion)
	}
	fmt.Printf("Downloading %s %s\n", packageName, packageVersion)
	downloadPackage(repoU.Path, packageVersion)
	if err := os.Mkdir("wyst_tmp", os.ModePerm); err != nil {
		ERR := fmt.Sprintf("%s", err)
		if !strings.Contains(strings.ToLower(ERR), "file exists") {
			panic(ERR)
		}
	}
	fmt.Printf("Extracting %s %s\n", packageName, packageVersion)
	Unzip("temp.zip", "wyst_tmp")
	entries, err := os.ReadDir("./wyst_tmp")
	if err != nil {
		panic(err)
	}
	if err := os.Mkdir("lib", os.ModePerm); err != nil {
		ERR := fmt.Sprintf("%s", err)
		if !strings.Contains(strings.ToLower(ERR), "file exists") {
			panic(ERR)
		}
	}
	for _, e := range entries {
		os.Rename("wyst_tmp/"+e.Name(), "lib/"+packageName)
	}
	if err := os.RemoveAll("wyst_tmp"); err != nil {
		panic(err)
	}
	if err := os.Remove("temp.zip"); err != nil {
		panic(err)
	}
	fmt.Printf("%s %s Installed\n", packageName, packageVersion)
}
