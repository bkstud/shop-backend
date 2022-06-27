package google

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shop/api/v1/auth"
	"shop/config"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var cred auth.Credentials
var conf *oauth2.Config

func init() {
	var err error
	cred, err = auth.ReadOauthSecrets("./secrets/google-creds.json", "GOOGLE")
	if err != nil {
		log.Panic("Failed to initialize google credentials")
	}
	redirectUrl := fmt.Sprintf("%s:%d/api/v1/auth/google", config.SERVER_ADDRESS, config.SERVER_PORT)
	conf = &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		RedirectURL:  redirectUrl,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			// You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
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
	client := auth.AuthHandler(c, conf)

	data := auth.GetUserData(c, client, "https://www.googleapis.com/oauth2/v3/userinfo")
	log.Println("google data:=", string(data))

	user := auth.Response{}
	if err := json.Unmarshal(data, &user); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error marshalling response. Please try agian."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "auth succeded"})

}
