package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"io/ioutil"
	"net/http"
	"reflect"
)

type GetWebRequest interface {
	FetchBytes(url string) []byte
}

type people struct {
	Number int `json:"number"`
}

type LiveGetWebRequest struct {
}

func (LiveGetWebRequest) FetchBytes(url string) []byte {
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}
	fmt.Println("spaceClient timeout :", spaceClient.Timeout, reflect.TypeOf(spaceClient.Timeout))

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return body
}
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func TestMyClient(client HttpClient) error {
	request, err := http.NewRequest("GET", "https://www.reallycoolurl.com", nil)
	if (err != nil) {
		return err
	}

	_, err = client.Do(request)
	if err != nil {
		return err
	}

	fmt.Println("Successful request.")
	return nil
}

type ClientMock struct {
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{}, nil
}

func GetAstronauts(getWebRequest GetWebRequest) int {
	url := "http://api.open-notify.org/astros.json"
	bodyBytes := getWebRequest.FetchBytes(url)
	peopleResult := people{}
	fmt.Println("PeopleResults : ", peopleResult)
	jsonErr := json.Unmarshal(bodyBytes, &peopleResult)
	fmt.Println("PeopleResults : ", jsonErr)
	fmt.Println("json.Unmarshal : ", json.Unmarshal(bodyBytes, &peopleResult))
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return peopleResult.Number
}

func main() {
	liveClient := LiveGetWebRequest{}
	fmt.Println("liveClient : ", liveClient)
	number := GetAstronauts(liveClient)

	fmt.Println("number : ", number)


	mock := &ClientMock{}
	err := TestMyClient(mock)
	if err != nil {
		fmt.Println(err.Error())
	}
}