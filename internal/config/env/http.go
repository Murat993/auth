package env

import (
	"github.com/Murat993/auth/internal/config"
	"net"
	"os"
	"time"

	"github.com/pkg/errors"
)

const (
	httpHostEnvName = "HTTP_HOST"
	httpPortEnvName = "HTTP_PORT"

	AuthPrefix = "Bearer "

	RefreshTokenSecretKey = "REFRESH_TOKEN_SECRET_KEY"
	AccessTokenSecretKey  = "ACCESS_TOKEN_SECRET_KEY"

	RefreshTokenExpiration = 60 * time.Minute
	AccessTokenExpiration  = 5 * time.Minute
)

type httpConfig struct {
	host string
	port string
}

func NewHTTPConfig() (config.HTTPConfig, error) {
	host := os.Getenv(httpHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("http host not found")
	}

	port := os.Getenv(httpPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("http port not found")
	}

	return &httpConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *httpConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
