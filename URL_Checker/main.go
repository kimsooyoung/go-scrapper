package main

import (
	"errors"
	"fmt"
	"net/http"
)

var urlErr = errors.New("URL Get Failed!!")

func hitURL(url string) error {
	fmt.Println("Checking : ", url)
	res, err := http.Get(url)

	if err != nil || res.StatusCode >= 400 {
		fmt.Println(err, res.StatusCode)
		return urlErr
	}
	return nil
}

func main() {
	var results = make(map[string]string)

	urls := []string{
		"https://www.naver.com/",
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.amazon.com/",
		"https://www.netflix.com/",
		"https://www.soundcloud.com/",
		"https://www.notion.com/",
		"https://www.helloworld.com/",
	}

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAIL"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}
