package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"flag"
)

var (
	url string
)

//go:generate moq -out tryingdependancyInjection3_moq_test.go . HttpClient

type HttpClient interface {
	Get(string) (*http.Response, error)
}

func init() {
	flag.StringVar(&url, "url", "http://google.com", "Which URL do we want to parse?")
	flag.Parse()
}

func main() {
	client := &http.Client{}
	err := send(client, url)

	if err != nil {
		panic(err)
	}
}

func send(client HttpClient, link string) error {
	response, err := client.Get(link)
	//fmt.Println(response)
	if err != nil {
		return err
	}

	if response == nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	fmt.Println(string(body))
	return nil
}
