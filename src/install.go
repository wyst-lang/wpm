package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Package struct {
	Author      string `json:"author"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type PackageIndex struct {
	Name string `json:"name"`
	Repo string `json:"repo"`
}

func installPackage(packageName, packageVersion string) {
	var pkgidx PackageIndex
	jsonBody := []byte(fmt.Sprintf(`{"name": "%s"}`, packageName))
	resBody := sendRequest(http.MethodGet, "http://localhost:3000", jsonBody)
	err := json.Unmarshal(resBody.Body, &pkgidx)
	if err != nil {
		fmt.Println("Json Error: ", err)
		return
	}

	fmt.Printf("repo: %s\n", pkgidx.Repo)

	repoU, err := url.Parse(pkgidx.Repo)
	if err != nil {
		panic(err)
	}

	if packageVersion == "" {
		packageVersion = getLatestVersion(repoU.Path)
	}
	downloadPackage(repoU.Path, packageVersion)
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
}
