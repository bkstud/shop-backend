package github

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shop/api/v1/auth"
	"shop/config"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var cred auth.Credentials
var conf *oauth2.Config

func init() {
	var err error
	cred, err = auth.ReadOauthSecrets("./secrets/github-creds.json", "GITHUB")
	if err != nil {
		log.Panic("Failed to initialize github credentials")
	}
	redirectUrl := fmt.Sprintf("%s:%d/api/v1/auth/github", config.SERVER_ADDRESS, config.SERVER_PORT)
	conf = &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		RedirectURL:  redirectUrl,
		Endpoint:     github.Endpoint,
		Scopes:       []string{"user", "public_profile", "user:email"},
	}
}
func LoginHandler(c *gin.Context) {
	auth.LoginHandler(c, conf)
}

func AuthHandler(c *gin.Context) {
	client := auth.AuthHandler(c, conf)
	data := auth.GetUserData(c, client, "https://api.github.com/user/emails")
	// TODO: This data is array of jsons we need to parse and one of them have field "primary": "true"
	// This one we use for user-id
	log.Println("github data:=", string(data))
	users := []auth.Response{}
	if err := json.Unmarshal(data, &users); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error marshalling response. Please try agian."})
		return
	}
	var primaryUser auth.Response
	for _, user := range users {
		if user.Primary {
			primaryUser = user
			break
		}
	}
	log.Println("primaryUser:=", primaryUser)
	c.JSON(http.StatusOK, gin.H{"message": "auth succeded"})
}
