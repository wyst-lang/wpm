package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Request struct {
	Body       []byte
	StatusCode int
}

type Package struct {
	Author      string `json:"author"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type PackageIndex struct {
	Name    string `json:"name"`
	Repo    string `json:"repo"`
	Message string `json:"message"`
}

func sendRequest(method, url string, jsonBody []byte) Request {
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		fmt.Printf("Error fetching package info: %v\n", err)
		return Request{Body: []byte(""), StatusCode: 402}
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	return Request{Body: resBody, StatusCode: res.StatusCode}
}

func getRepoURL(packageName string) (string, error) {
	var pkgidx PackageIndex
	jsonBody := []byte(fmt.Sprintf(`{"name": "%s"}`, packageName))
	req := sendRequest(http.MethodGet, "http://localhost:3000", jsonBody)
	err := json.Unmarshal(req.Body, &pkgidx)
	if err != nil {
		fmt.Println("Json Error: ", err)
		return "", err
	}
	if req.StatusCode != 200 {
		return "", fmt.Errorf(pkgidx.Message)
	}
	return pkgidx.Repo, nil
}
