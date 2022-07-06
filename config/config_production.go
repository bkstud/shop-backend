//go:build production
// +build production

package config

import "os"

var (
	SERVER_ADDRESS    = "https://" + os.Getenv("WEBSITE_HOSTNAME")
	FRONTEND_HOSTNAME = os.Getenv("FRONTEND_HOSTNAME")
	FRONTEND_ADDRESS  = "http://" + os.Getenv("FRONTEND_HOSTNAME")
)

const (
	SERVER_PORT = 443
	ENV         = "PRODUCTION"
)
