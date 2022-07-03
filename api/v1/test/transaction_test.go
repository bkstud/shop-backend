//go:build test
// +build test

package test

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"shop/api/v1/model"
	"shop/config"
	"testing"
)

var ENDPOINT = struct {
	TRANSACTION string
	ITEM        string
	API         string
}{
	TRANSACTION: fmt.Sprintf("%s:%d/api/v1/transactions", config.SERVER_ADDRESS, config.SERVER_PORT),
	ITEM:        fmt.Sprintf("%s:%d/api/v1/items", config.SERVER_ADDRESS, config.SERVER_PORT),
	API:         fmt.Sprintf("%s:%d/api/v1/", config.SERVER_ADDRESS, config.SERVER_PORT),
}

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}
func TestGetEndpoint(t *testing.T) {

	resp, err := http.Get(ENDPOINT.TRANSACTION)
	if err != nil {
		t.Errorf("Error in get response %s", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Got response %d instead 200", resp.StatusCode)
	}
	var outArr []model.Transaction
	json.NewDecoder(resp.Body).Decode(&outArr)
	if len(outArr) > 0 {
		t.Errorf("Error in response got %v instead of []", outArr)
	}

}
func TestCreateTransaction(t *testing.T) {
	item := model.Item{
		Name:        "Nike Shoes",
		Description: "Running shoes",
		Status:      "available",
		Price:       180.80,
	}

	user := model.User{
		Email: "test@gmail.com",
		Name:  "Adam",
		Type:  "google",
	}

	transaction := model.Transaction{
		User:   user,
		Item:   item,
		Status: "pending",
	}
	data, _ := json.Marshal(transaction)

	resp, err := http.Post(ENDPOINT.TRANSACTION, "application/json",
		bytes.NewBuffer(data))

	if err != nil {
		t.Errorf("Error in post response %s", err)
	}
	fmt.Println(resp)
}
