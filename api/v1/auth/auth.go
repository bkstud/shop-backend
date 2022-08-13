package auth

import (
	"context"
	"fmt"
	"io/ioutil"
	"shop/api/v1/controller"
	"shop/api/v1/database"
	"shop/api/v1/model"
	"shop/api/v1/utils"
	"shop/config"

	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// Warning this may not be thread safe and may not work in case of heavy load
// left for demo purposes
var LastLocation string

func LoginHandler(c *gin.Context, conf *oauth2.Config) {
	LastLocation = c.DefaultQuery("location", "")
	state, err := utils.RandString(32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	session := sessions.Default(c)
	session.Set("state", state)
	fmt.Println(c.Request.URL.User)
	fmt.Println(c.Request.Referer())

	err = session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error while saving session."})
		return
	}

	link := conf.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, link)
}

func AuthHandler(c *gin.Context, conf *oauth2.Config) *http.Client {
	session := sessions.Default(c)
	retrievedState := session.Get("state")
	queryState := c.Request.URL.Query().Get("state")
	if retrievedState != queryState {
		log.Printf("Invalid session state: retrieved: %s; Param: %s", retrievedState, queryState)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid session state."})
		return nil
	}

	code := c.Request.URL.Query().Get("code")
	tok, err := conf.Exchange(context.Background(), code)
	if err != nil {
		log.Println("error:=", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Login failed. Please try again."})
		return nil
	}

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
	newToken, err := controller.CreateOrUpdateToken(email)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while saving session. Please try again."})
		return
	}
	c.Set("bearer", newToken.Bearer)
}

// Creates new user if does not exists with following email or returns existent one.
func CreateUserFromResponse(responseUser *Response) model.User {
	user := model.User{}
	db := database.Database
	exists := db.First(&user, "email = ?", responseUser.Email)
	if exists.Error != nil {
		user.Email = responseUser.Email
		user.Name = responseUser.Name
		user.Type = responseUser.Type
		db.Create(&user)
	}
	return user
}

func RedirectBack(c *gin.Context) {
	token := c.MustGet("bearer")
	endpoint := fmt.Sprintf("/login/?location=%s&token=%s", LastLocation, token)
	c.Redirect(http.StatusTemporaryRedirect, config.FRONTEND_ADDRESS+endpoint)
}
