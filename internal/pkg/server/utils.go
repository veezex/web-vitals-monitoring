package server

import "github.com/veezex/web-vitals-monitoring/server/internal/pkg/config"

func protocol(config config.Config) string {
	protocol := "https"
	if !config.GetUseHttps() {
		protocol = "http"
	}
	return protocol
}
