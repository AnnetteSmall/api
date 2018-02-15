package fib

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	client := http.Client{}
	response, err := client.Get("http://google.com")

	if err != nil {
		panic(err)
	}

	if response == nil {
		fmt.Println("Received an empty response!")
		return
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Couldn't read body of response!")
		return
	}

	fmt.Println(string(body))
}
