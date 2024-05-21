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
	progress := ProgressBar{total: 10, length: 20, enabled: true}
	progress.change(1, "", "Fetching")
	pkgidx, err := getPackage(packageName)
	if err != nil {
		fmt.Printf("Error getting repo URL: %s\n", err)
		return
	}
	repoU, err := url.Parse(pkgidx.Repo)
	if err != nil {
		fmt.Printf("Error parsing repo URL: %s\n", err)
	}
	if packageVersion == "" {
		packageVersion = pkgidx.Latest
		progress.change(3, "", fmt.Sprintf("Using version %s", packageVersion))
	}
	progress.change(6, "", "Downloading")
	downloadPackage(repoU.Path, packageVersion)
	if err := os.Mkdir("wyst_tmp", os.ModePerm); err != nil {
		ERR := fmt.Sprintf("%s", err)
		if !strings.Contains(strings.ToLower(ERR), "file exists") {
			panic(ERR)
		}
	}
	progress.change(9, "", "Extracting")
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
	progress.change(10, "", "Installed")
	progress.clean()
	fmt.Printf("%s %s Installed\n\n", packageName, packageVersion)
}
