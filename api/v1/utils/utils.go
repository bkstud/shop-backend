package utils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

func RandString(l int) (string, error) {
	b := make([]byte, l)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// Credentials which stores google ids.
type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

func ReadOauthSecrets(secretfile string, varPostfix string) (Credentials, error) {
	var cred Credentials
	if _, err := os.Stat(secretfile); errors.Is(err, os.ErrNotExist) {
		// If config file no present read from ENV
		varid := "VAR_ID_" + varPostfix
		cid, exists := os.LookupEnv(varid)
		if !exists {
			log.Printf("Variable '%s' not set\n", varid)
			return cred, errors.New("variable does not exists: " + varid)
		}
		cred.Cid = cid

		varsecret := "VAR_SECRET_" + varPostfix
		cid, exists = os.LookupEnv(varsecret)
		if !exists {
			log.Printf("Variable '%s' not set\n", varsecret)
			return cred, errors.New("variable does not exists: " + varsecret)
		}
		cred.Csecret = cid

	} else {
		file, err := ioutil.ReadFile(secretfile)
		if err != nil {
			log.Printf("File error: %v\n", err)
			return cred, err
		}
		if err := json.Unmarshal(file, &cred); err != nil {
			log.Println("unable to marshal data")
			return cred, err
		}
	}
	return cred, nil
}
