package main

import (
	"testing"
	"net/http/httptest"
	"net/url"
	"fmt"
)

func TestSafeName(t *testing.T){
	itemName := "apple"
	appID := "d3a8bee3"
	appKey := "fb2b5f577a67a7f542ec03297b29b372"
	safeItemName := url.QueryEscape(itemName)
	url := fmt.Sprintf("http://api.nutritionix.com/v1_1/search/%s?results=0:10&fields=nf_total_carbohydrate,nf_sugars,nf_dietary_fiber,item_name,brand_name,item_id,nf_calories&appId=%s&appKey=%s", safeItemName, appID, appKey)

	req, err := httptest.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("NewRequest : ", err)
		t.Fatal("NewRequest: ", err)
		return
	}
}
