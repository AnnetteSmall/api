package main

import "testing"

type testWebRequest struct {
}

func (testWebRequest) FetchBytes(url string) []byte {
	return []byte(`{"number": 2}`)
}

func TestGetAstronauts(t *testing.T) {
	amount := GetAstronauts(testWebRequest{})
	if amount != 1 {
		t.Errorf("People in space, got: %d, want: %d.", amount, 1)
	}
}

