package github

import (
	"fmt"
	"shop/api/v1/auth"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var cred auth.Credentials
var conf *oauth2.Config

func init() {
	var err error
	cred, err = auth.ReadOauthSecrets("./github-creds.json", "GITHUB")
	if err != nil {
		fmt.Println("Failed to initialize github credentials")
	}

	conf = &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		RedirectURL:  "http://localhost:80/auth/github",
		Endpoint:     github.Endpoint,
	}
}
func LoginHandler(c *gin.Context) {
	auth.LoginHandler(c, conf)
}

func AuthHandler(c *gin.Context) {
	auth.AuthHandler(c, conf)
}
