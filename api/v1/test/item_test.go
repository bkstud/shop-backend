package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"shop/api/v1/model"
	"testing"
)

var ENDPOINT = "localhost:5000/api/v1/items"

func TestGetEndpoint(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("http://%s", ENDPOINT))
	if err != nil {
		t.Errorf("Error in get response %s", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Got response %d instead 200", resp.StatusCode)
	}
	var outArr []model.Item
	json.NewDecoder(resp.Body).Decode(&outArr)
	if len(outArr) > 0 {
		t.Errorf("Error in response got %v instead of []", outArr)
	}

}

func TestCreateItem(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Name":        "Nike Shoes",
		"Description": "Running shoes",
	})

	resp, err := http.Post(fmt.Sprintf("http://%s", ENDPOINT), "application/json",
		bytes.NewBuffer(postBody))

	if err != nil {
		t.Errorf("Error in get response %s", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Got response %d instead 200", resp.StatusCode)
	}

	var out = model.Item{}
	json.NewDecoder(resp.Body).Decode(&out)
	if out.Name != "Nike Shoes" {
		t.Errorf("Response have wrong name '%s' should be 'Nike Shoes'", out.Name)
	}
	if out.Description != "Running shoes" {
		t.Errorf("Response have wrong description '%s' should be 'Running shoes'", out.Name)
	}
}

func TestEditItem(t *testing.T) {
	resp, _ := http.Get(fmt.Sprintf("http://%s", ENDPOINT))
	if resp.StatusCode != 200 {
		t.Errorf("Got response %d instead 200", resp.StatusCode)
	}

	var outArr []model.Item
	json.NewDecoder(resp.Body).Decode(&outArr)
	id := outArr[0].ID

	patchBody, _ := json.Marshal(map[string]string{
		"Name":        "Adidas shoes",
		"Description": "Sneakers shoes",
	})

	req, _ := http.NewRequest(http.MethodPatch,
		fmt.Sprintf("http://%s/%d", ENDPOINT, id),
		bytes.NewBuffer(patchBody))
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("PATCH request failed with error %v", err)
	}
	if res == nil {
		t.Errorf("Failed to get response from %v", req)
	}

	if res == nil || res.StatusCode != 200 {
		t.Errorf("Got response %d instead 200", res.StatusCode)
	}
	var patchOut model.Item
	json.NewDecoder(res.Body).Decode(&patchOut)

	if patchOut.Name != "Adidas shoes" {
		t.Errorf("Response have wrong name '%s' should be 'Adidas shoes'", patchOut.Name)
	}
	if patchOut.Description != "Sneakers shoes" {
		t.Errorf("Response have wrong description '%s' should be 'Sneakers shoes'", patchOut.Description)
	}
}
func TestRemoveItem(t *testing.T) {
	resp, _ := http.Get(fmt.Sprintf("http://%s", ENDPOINT))
	if resp.StatusCode != 200 {
		t.Errorf("Got response %d instead 200", resp.StatusCode)
	}

	var outArr []model.Item
	json.NewDecoder(resp.Body).Decode(&outArr)
	if len(outArr) > 1 {
		t.Errorf("Error in response got %v instead of single value", outArr)
	}
	id := outArr[0].ID

	req, _ := http.NewRequest(http.MethodDelete,
		fmt.Sprintf("http://%s/%d", ENDPOINT, id),
		nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("DELETE request failed with error %v", err)
	}
	if res == nil {
		t.Errorf("Failed to get response from %v", req)
	}
	// Recall first test to check if no items are available now
	TestGetEndpoint(t)
}
