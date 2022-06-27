package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"shop/api/v1/model"

	"log"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

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

func RandToken(l int) (string, error) {
	b := make([]byte, l)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func LoginHandler(c *gin.Context, conf *oauth2.Config) {
	log.Println("login handler")
	state, err := RandToken(32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	session := sessions.Default(c)
	session.Set("state", state)
	err = session.Save()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{"message": "Error while saving session."})
		return
	}

	link := conf.AuthCodeURL(state)
	log.Printf("link:= %s", link)
	c.Redirect(http.StatusTemporaryRedirect, link)
}

func AuthHandler(c *gin.Context, conf *oauth2.Config) *http.Client {
	session := sessions.Default(c)
	retrievedState := session.Get("state")
	queryState := c.Request.URL.Query().Get("state")
	if retrievedState != queryState {
		log.Printf("Invalid session state: retrieved: %s; Param: %s", retrievedState, queryState)
		c.HTML(http.StatusUnauthorized, "error.tmpl", gin.H{"message": "Invalid session state."})
		return nil
	}

	code := c.Request.URL.Query().Get("code")
	log.Printf("code:= %v", code)
	tok, err := conf.Exchange(context.Background(), code)
	if err != nil {
		log.Println("error:=", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Login failed. Please try again."})
		return nil
	}
	// sessionToken, _ := RandToken(32)
	client := conf.Client(context.Background(), tok)

	return client
}

func GetUserData(c *gin.Context, client *http.Client, dataEndpoint string) []byte {
	userinfo, err := client.Get(dataEndpoint)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return nil
	}
	defer userinfo.Body.Close()
	data, _ := ioutil.ReadAll(userinfo.Body)
	return data
}

func SetIdentityEmail(c *gin.Context, email string) {
	session := sessions.Default(c)
	u := model.User{}
	session.Set("user-id", email)
	err := session.Save()
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Error while saving session. Please try again."})
		return
	}
	u.Email = email
}
