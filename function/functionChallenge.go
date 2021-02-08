package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ctype := resp.Header.Get("Content-Length")
	if ctype == "" {
		return "", fmt.Errorf("cant find  the return error")
	}

	return ctype, nil

}

func main() {

	ctype, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("ERROR : %s\n", err)
	} else {
		fmt.Printf(ctype)
	}

}
