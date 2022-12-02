package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	//sendRequest()
	httpClient()
}

func sendRequest() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

func httpClient() {
	client := http.Client{
		Timeout: 6 * time.Second,
	}
	resp, err := client.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
