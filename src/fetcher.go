package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	Latest  string `json:"latest"`
	Message string `json:"message"`
}

type Message struct {
	Message string `json:"message"`
}

func sendRequest(method, url string, jsonBody []byte) (Request, error) {
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return Request{Body: []byte(""), StatusCode: 400}, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Request{Body: []byte(""), StatusCode: 400}, err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Request{Body: []byte(""), StatusCode: 400}, err
	}
	return Request{Body: resBody, StatusCode: res.StatusCode}, nil
}

func getPackage(packageName string) (PackageIndex, error) {
	var pkgidx PackageIndex
	jsonBody := []byte(fmt.Sprintf(`{"name": "%s"}`, packageName))
	req, err := sendRequest(http.MethodGet, URL, jsonBody)
	if err != nil {
		fmt.Printf("Fetching error: %s\n", err)
	}
	err = json.Unmarshal(req.Body, &pkgidx)
	if err != nil {
		fmt.Println("Json Error: ", err)
		return pkgidx, err
	}
	if req.StatusCode != 200 {
		return pkgidx, fmt.Errorf(pkgidx.Message)
	}
	return pkgidx, nil
}

func getMessage(req Request) error {
	var msg Message
	err := json.Unmarshal(req.Body, &msg)
	if err != nil {
		return fmt.Errorf("JsonError: %v", err)
	}
	if msg.Message != "ok" {
		return fmt.Errorf(msg.Message)
	}
	return nil
}
