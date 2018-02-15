package main

import (
	"net/url"
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

type nutrition struct {
	TotalHits int `json:"total_hits"`
	MaxScore float64 `json:"max_score"`
	Hits []struct {
		Index string `json:"_index"`
		Type string `json:"_type"`
		ID string `json:"_id"`
		Score float64 `json:"_score"`
		Fields struct {
			ItemID string `json:"item_id"`
			ItemName string `json:"item_name"`
			BrandName string `json:"brand_name"`
			NfCalories float64 `json:"nf_calories"`
			NfTotalCarbohydrate float64 `json:"nf_total_carbohydrate"`
			NfDietaryFiber float64 `json:"nf_dietary_fiber"`
			NfSugars float64 `json:"nf_sugars"`
			NfServingSizeQty int `json:"nf_serving_size_qty"`
			NfServingSizeUnit string `json:"nf_serving_size_unit"`
		} `json:"fields"`
	} `json:"hits"`
}

func main()  {
	itemName := "apple"
	appID := "d3a8bee3"
	appKey := "fb2b5f577a67a7f542ec03297b29b372"
	safeItemName := url.QueryEscape(itemName)

	url := fmt.Sprintf("http://api.nutritionix.com/v1_1/search/%s?results=0:10&fields=nf_total_carbohydrate,nf_sugars,nf_dietary_fiber,item_name,brand_name,item_id,nf_calories&appId=%s&appKey=%s", safeItemName, appID, appKey)
	//fmt.Println("URL :", url)
	//Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln("NewRequest : ", err)
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}
	err := send(client, url)

	if err != nil {
		panic(err)
	}
	fmt.Println("client = ", client)

	resp, err:= client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	defer resp.Body.Close()

	var record nutrition

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	num := len(record.Hits)

	for i:=0; i < num ; i++  {
		fmt.Println("Item : ", record.Hits[i].Fields.ItemName)
		fmt.Println("Calories :", record.Hits[i].Fields.NfCalories)
		fmt.Println("Net Carbs : ", record.Hits[i].Fields.NfTotalCarbohydrate - record.Hits[i].Fields.NfDietaryFiber - record.Hits[i].Fields.NfSugars )
		fmt.Println("***")
	}
	
	}

func send(client HttpClient, link string) error {
	response, err := client.Do(link)
		return err
	}
	if response == nil {
		return err
	}
	var record nutrition

	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		log.Println(err)
	}

}