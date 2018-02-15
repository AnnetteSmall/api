package main

import (
	"net/url"
	"net/http"
	"fmt"
	"log"
)
type HttpClient interface {
	Get(string) (*http.Response, error)
}


func main(){

}

func safename() string  {
	i := url.QueryEscape("sugar cube")
	return i
}

func makeurl () string {
	url := fmt.Sprintf("http://api.nutritionix.com/v1_1/search/%s?results=0:10&fields=nf_total_carbohydrate,nf_sugars,nf_dietary_fiber,item_name,brand_name,item_id,nf_calories&appId=%s&appKey=%s", safename(), "d3a8bee3", "fb2b5f577a67a7f542ec03297b29b372")
	return url
}

func getnutrition (nutrition struct{}) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln("NewRequest : ", err)
		log.Fatal("NewRequest: ", err)
		return
	}

}