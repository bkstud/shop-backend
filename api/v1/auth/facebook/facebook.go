package facebook

import (
	"encoding/json"
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
		Scopes:       []string{"email", "public_profile"},
		Endpoint:     facebook.Endpoint,
	}
}

// LoginHandler handles the login procedure.
func LoginHandler(c *gin.Context) {
	auth.LoginHandler(c, conf)
}

func AuthHandler(c *gin.Context) {
	client := auth.AuthHandler(c, conf)
	if client == nil {
		return
	}
	data := auth.GetUserData(c, client, "https://graph.facebook.com/me?fields=name,email")
	if data == nil {
		return
	}

	resp := auth.Response{}
	if err := json.Unmarshal(data, &resp); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error marshalling response. Please try agian."})
		return
	}
	resp.Type = "facebook"
	auth.SetIdentityEmail(c, resp.Email)
	auth.CreateUserFromResponse(&resp)
	c.Redirect(http.StatusTemporaryRedirect, config.FRONTEND_ADDRESS+auth.LastLocation)
}
