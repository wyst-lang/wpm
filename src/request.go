package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Request struct {
	Body       []byte
	StatusCode int
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
