package server

import (
	"fmt"
	"github.com/veezex/web-vitals-monitoring/server/internal/pkg/config"
	"github.com/veezex/web-vitals-monitoring/server/internal/pkg/db"
	"net/http"
	"strconv"
)

func RunServer(config config.Config, db db.DB) error {
	http.HandleFunc("/metric", createHandleMetric(db))
	http.HandleFunc("/script", createHandleScript(config))

	fmt.Println("Server is running on port:", config.GetPort())
	fmt.Println("Client script:")
	fmt.Printf(scriptExample(config))

	if config.GetUseHttps() {
		return http.ListenAndServeTLS(":"+strconv.Itoa(config.GetPort()), "fullchain.pem", "privkey.pem", nil)
	} else {
		return http.ListenAndServe(":"+strconv.Itoa(config.GetPort()), nil)
	}
}

func scriptExample(config config.Config) string {
	return fmt.Sprintf("<script type=\"module\" async src=\"%s://%s:%d/script\"></script>\n", protocol(config), config.GetDomain(), config.GetPort())
}
