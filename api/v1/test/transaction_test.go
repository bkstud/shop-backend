//go:build test
// +build test

package test

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
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
func TestGetNoTransactions(t *testing.T) {

	resp, err := http.Get(ENDPOINT.TRANSACTION)
	if err != nil {
		t.Errorf("Error in get response %s", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Got response %d instead 200", resp.StatusCode)
	}
	var outArr []model.Transaction
	json.NewDecoder(resp.Body).Decode(&outArr)
	assert.Empty(t, outArr, "Array should be empty")

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

	if resp.StatusCode != 200 {
		t.Errorf("Got response %d instead 200", resp.StatusCode)
		data, _ := io.ReadAll(resp.Body)
		t.Errorf("Response json: %s", string(data))
	}
	var out = model.Transaction{}
	json.NewDecoder(resp.Body).Decode(&out)
	assert.Equal(t, out.User.Name, transaction.User.Name, "Names should be the same")
	assert.Equal(t, out.User.Email, transaction.User.Email, "Emails should be the same")
	assert.Equal(t, out.Item.Name, transaction.Item.Name, "Names should be the same")
	assert.Equal(t, out.Item.Description, transaction.Item.Description, "Descriptions should be the same")
}

func TestEditTransaction(t *testing.T) {
	transaction := model.Transaction{
		Status: "finished",
	}
	data, _ := json.Marshal(transaction)
	id := 1
	req, _ := http.NewRequest(http.MethodPatch,
		fmt.Sprintf("%s/%d", ENDPOINT.TRANSACTION, id),
		bytes.NewBuffer(data))
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("PATCH request failed with error %v", err)
	}
	if res == nil {
		t.Errorf("Failed to get response from %v", req)
	}

	if res.StatusCode != 200 {
		t.Errorf("Got response %d instead 200", res.StatusCode)
	}
	var patchOut model.Transaction
	json.NewDecoder(res.Body).Decode(&patchOut)
	assert.Equal(t, patchOut.Status, transaction.Status, "Statuses should be the same")
}

func TestDeleteTransaction(t *testing.T) {
	id := 1
	req, _ := http.NewRequest(http.MethodDelete,
		fmt.Sprintf("%s/%d", ENDPOINT.TRANSACTION, id),
		nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("DELETE request failed with error %v", err)
	}
	if res == nil {
		t.Errorf("Failed to get response from %v", req)
	}

	// DELETE created item
	req, _ = http.NewRequest(http.MethodDelete,
		fmt.Sprintf("%s/%d", ENDPOINT.ITEM, id),
		nil)
	res, err = http.DefaultClient.Do(req)

	// Recall first test to check if no items are available now
	TestGetNoTransactions(t)

}
