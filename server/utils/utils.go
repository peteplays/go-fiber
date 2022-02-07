package utils

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

func CreateRandomNumber() int {
	min := 1
	max := 82
	return rand.Intn(max-min) + min
}

func HttpGet(path string) *http.Response {
	var baseUrl string
	if !strings.Contains(path, "http") {
		baseUrl = os.Getenv("SW_API_ENDPOINT")
	}

	fmt.Println("URL:---", baseUrl+path)

	res, err := http.Get(baseUrl + path)
	if err != nil {
		log.Fatal(err)
	}

	// defer resp.Body.Close()
	return res
}

func GetURLParams(urls []string) []string {
	var splitUrls []string
	for _, url := range urls {
		urlSplit := strings.Split(url, "/")
		splitUrls = append(splitUrls, fmt.Sprintf("/%s/%s", urlSplit[4], urlSplit[5]))
	}

	return splitUrls
}
