package server

import "github.com/veezex/web-vitals-monitoring/server/internal/pkg/config"
import "net/http"

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "content-type")
}

func protocol(config config.Config) string {
	protocol := "https"
	if !config.GetUseHttps() {
		protocol = "http"
	}
	return protocol
}
