package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

type PackageIndex struct {
	Versions map[string]PackageVersion `json:"versions"`
}

func installPackage(packageName, packageVersion string) {
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

	if packageVersion == "" {
		packageVersion = getLatestVersion(*pkgIndex)
	}

	pkgVersion, exists := pkgIndex.Versions[packageVersion]
	if !exists {
		fmt.Printf("Version %s not found for package %s\n", packageVersion, packageName)
		return
	}

	downloadAndSavePackage(pkgVersion.URL, packageName, packageVersion)
}

func getRepoURL(packageName string) (string, error) {
	url := "http://localhost:3000"
	requestBody, err := json.Marshal(map[string]string{
		"name": packageName,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	var repoResponse RepoResponse
	if err := json.NewDecoder(resp.Body).Decode(&repoResponse); err != nil {
		return "", err
	}

	if repoResponse.Message != "ok" {
		return "", fmt.Errorf("unexpected message: %s", repoResponse.Message)
	}

	return repoResponse.Repo, nil
}

func findIndexURL(repoURL string) (string, error) {
	possibleURLs := []string{
		repoURL,
		repoURL + "/index.json",
		repoURL + "/versions.json",
	}

	for _, url := range possibleURLs {
		resp, err := http.Get(url)
		if err != nil {
			continue
		}
		if resp.StatusCode == http.StatusOK {
			resp.Body.Close()
			return url, nil
		}
		resp.Body.Close()
	}

	return "", fmt.Errorf("no valid index URL found at %s", repoURL)
}

func fetchPackageIndex(url string) (*PackageIndex, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching package index: %s", resp.Status)
	}

	var pkgIndex PackageIndex
	if err := json.NewDecoder(resp.Body).Decode(&pkgIndex); err != nil {
		return nil, err
	}

	return &pkgIndex, nil
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

	// Infer filename from URL
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
