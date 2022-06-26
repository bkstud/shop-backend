package facebook

import (
	"fmt"
	"log"
	"net/http"
	"shop/api/v1/auth"
	"shop/config"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var cred auth.Credentials
var conf *oauth2.Config

func init() {
	var err error
	cred, err = auth.ReadOauthSecrets("./secrets/facebook-creds.json", "FACEBOOK")
	if err != nil {
		log.Panic("Failed to initialize facebook credentials")
	}
	redirectUrl := fmt.Sprintf("%s:%d/api/v1/auth/facebook", config.SERVER_ADDRESS, config.SERVER_PORT)
	conf = &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		RedirectURL:  redirectUrl,
		Scopes:       []string{"email", "read_stream"},
		Endpoint:     facebook.Endpoint,
	}
}

// LoginHandler handles the login procedure.
func LoginHandler(c *gin.Context) {
	log.Println("facebook login handler")
	auth.LoginHandler(c, conf)
}

func AuthHandler(c *gin.Context) {
	log.Printf("facebook auth handler")
	client := auth.AuthHandler(c, conf)
	if client == nil {
		return
	}
	data := auth.GetUserData(c, client, "https://graph.facebook.com/me?fields=name,email")
	if data == nil {
		return
	}
	log.Println("fb data:=", string(data))

	c.JSON(http.StatusOK, gin.H{"message": "auth succeded"})
}
