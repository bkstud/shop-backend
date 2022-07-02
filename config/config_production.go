//go:build production
// +build production

package config

import "os"

var (
	SERVER_ADDRESS = "https://" + os.Getenv("WEBSITE_HOSTNAME")
)

const (
	SERVER_PORT = 443
	ENV         = "PRODUCTION"
)
