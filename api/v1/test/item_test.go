//go:build test
// +build test

package test

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"shop/api/v1/model"
	"testing"
)

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}
func TestGetNoItems(t *testing.T) {

	resp, err := http.Get(ENDPOINT.ITEM)
	if err != nil {
		t.Errorf("Error in get response %s", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf(not_200_response, resp.StatusCode)
	}
	var outArr []model.Item
	json.NewDecoder(resp.Body).Decode(&outArr)
	if len(outArr) > 0 {
		t.Errorf("Error in response got %v instead of []", outArr)
	}

}

func TestCreateItem(t *testing.T) {
	postBody, _ := json.Marshal(model.Item{
		Name:        "Nike Shoes",
		Description: "Running shoes",
		Status:      "available",
		Price:       180.80,
	})
	resp, err := http.Post(ENDPOINT.ITEM, "application/json",
		bytes.NewBuffer(postBody))

	if err != nil {
		t.Errorf("Error in post response %s", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf(not_200_response, resp.StatusCode)
		data, _ := io.ReadAll(resp.Body)
		t.Errorf("Response json: %s", string(data))
	}

	var out = model.Item{}
	json.NewDecoder(resp.Body).Decode(&out)
	if out.Name != "Nike Shoes" {
		t.Errorf("Response have wrong name '%s' should be 'Nike Shoes'", out.Name)
	}
	if out.Description != "Running shoes" {
		t.Errorf("Response have wrong description '%s' should be 'Running shoes'", out.Name)
	}
	if out.Status != "available" {
		t.Errorf("Response have wrong description '%s' should be 'sold'", out.Status)
	}
}

func TestEditItem(t *testing.T) {
	resp, _ := http.Get(ENDPOINT.ITEM)
	if resp.StatusCode != 200 {
		t.Errorf(not_200_response, resp.StatusCode)
	}

	var outArr []model.Item
	json.NewDecoder(resp.Body).Decode(&outArr)
	id := outArr[0].ID

	patchBody, _ := json.Marshal(model.Item{
		Name:        "Adidas shoes",
		Description: "Sneakers shoes",
		Status:      "sold",
		Price:       220.22,
	})

	req, _ := http.NewRequest(http.MethodPatch,
		fmt.Sprintf("%s/%d", ENDPOINT.ITEM, id),
		bytes.NewBuffer(patchBody))
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("PATCH request failed with error %v", err)
	}
	if res == nil {
		t.Errorf("Failed to get response from %v", req)
	}

	if res.StatusCode != 200 {
		t.Errorf(not_200_response, res.StatusCode)
	}
	var patchOut model.Item
	json.NewDecoder(res.Body).Decode(&patchOut)

	if patchOut.Name != "Adidas shoes" {
		t.Errorf("Response have wrong name '%s' should be 'Adidas shoes'", patchOut.Name)
	}
	if patchOut.Description != "Sneakers shoes" {
		t.Errorf("Response have wrong description '%s' should be 'Sneakers shoes'", patchOut.Description)
	}
	if patchOut.Status != "sold" {
		t.Errorf("Response have wrong description '%s' should be 'sold'", patchOut.Status)
	}
}
func TestRemoveItem(t *testing.T) {
	resp, _ := http.Get(ENDPOINT.ITEM)
	if resp.StatusCode != 200 {
		t.Errorf(not_200_response, resp.StatusCode)
	}

	var outArr []model.Item
	json.NewDecoder(resp.Body).Decode(&outArr)
	if len(outArr) > 1 {
		t.Errorf("Error in response got %v instead of single value", outArr)
	}
	id := outArr[0].ID

	req, _ := http.NewRequest(http.MethodDelete,
		fmt.Sprintf("%s/%d", ENDPOINT.ITEM, id),
		nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("DELETE request failed with error %v", err)
	}
	if res == nil {
		t.Errorf("Failed to get response from %v", req)
	}
	// Recall first test to check if no items are available now
	TestGetNoItems(t)
}
