package google

import (
	"fmt"
	"log"
	"shop/api/v1/auth"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var cred auth.Credentials
var conf *oauth2.Config

func init() {
	var err error
	cred, err = auth.ReadOauthSecrets("./google-creds.json", "GOOGLE")
	if err != nil {
		fmt.Println("Failed to initialize google credentials")
	}

	conf = &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		RedirectURL:  "http://127.0.0.1:80/auth/google",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		},
		Endpoint: google.Endpoint,
	}
}

// LoginHandler handles the login procedure.
func LoginHandler(c *gin.Context) {
	log.Println("google login handler")
	auth.LoginHandler(c, conf)
}

func AuthHandler(c *gin.Context) {
	log.Printf("google auth handler")
	auth.AuthHandler(c, conf)
}
