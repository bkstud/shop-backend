package github

import (
	"fmt"
	"log"
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
	}
}
func LoginHandler(c *gin.Context) {
	auth.LoginHandler(c, conf)
}

func AuthHandler(c *gin.Context) {
	auth.AuthHandler(c, conf)
}
