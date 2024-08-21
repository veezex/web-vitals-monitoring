package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config interface {
	GetPort() int
	GetDomain() string
	GetUseHttps() bool
}

type configImpl struct {
	domain   string
	port     int
	useHttps bool
}

func New() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	domainString := os.Getenv("DOMAIN")
	portString := os.Getenv("PORT")
	useHttpsString := os.Getenv("USE_HTTPS")

	port, err := strconv.Atoi(portString)
	if err != nil {
		return nil, err
	}

	return &configImpl{
		domain:   domainString,
		useHttps: useHttpsString == "1",
		port:     port,
	}, nil
}

func (c *configImpl) GetPort() int {
	return c.port
}

func (c *configImpl) GetDomain() string {
	return c.domain
}

func (c *configImpl) GetUseHttps() bool {
	return c.useHttps
}
