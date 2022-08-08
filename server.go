package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"shop/api/v1/auth"
	"shop/api/v1/database"
	"shop/api/v1/route"
	"shop/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
)

var (
	routerHttps = gin.Default()
)

func main() {

	token, err := auth.RandToken(64)
	if err != nil {
		log.Fatal("unable to generate random token: ", err)
	}
	store := gormsessions.NewStore(database.Database, true, []byte(token))
	// cookie.NewStore([]byte(token))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
	})
	corsConf := cors.DefaultConfig()
	corsConf.AllowOrigins = []string{"http://" + config.FRONTEND_HOSTNAME,
		"https://" + config.FRONTEND_HOSTNAME}
	corsConf.AllowCredentials = true
	routerHttps.Use(cors.New(corsConf))

	routerHttps.Use(gin.Recovery())
	routerHttps.Use(sessions.Sessions("store", store))

	v1 := routerHttps.Group("/api/v1")
	route.AddRoutes(v1)

	if config.ENV == "PRODUCTION" {
		routerHttps.Run(":80")
	} else {
		GetCertAndKey()
		routerHttps.RunTLS(fmt.Sprintf(":%d", config.SERVER_PORT), "./cert/cert.pem", "./cert/key.pem")
	}

}

func GetCertAndKey() {
	varFile := map[string]string{"VAR_CERT": "./cert/cert.pem", "VAR_PRIVKEY": "./cert/key.pem"}
	if _, err := os.Stat("./cert"); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir("./cert", 0755)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		return
	}

	for variable, file := range varFile {
		value, exists := os.LookupEnv(variable)
		if exists {
			err := ioutil.WriteFile(file, []byte(value), 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
